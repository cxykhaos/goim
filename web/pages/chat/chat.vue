<template>
	<view class="content">
		<view class="content-box" @touchstart="touchstart" id="content-box" :class="{'content-showfn':showFunBtn}">
			<!-- 背景图- 定位方式 -->
			<image class="content-box-bg" :src="_user_info.chatBgImg" :style="{ height: imgHeight }"></image>
			<view class="content-box-loading" v-if="!loading">
				<u-loading mode="flower"></u-loading>
			</view>
			<view class="message" v-for="(item, index) in messageList" :key="index" :id="`msg-${item.hasBeenSentId}`">
				<view class="message-item " :class="item.isItMe ? 'right' : 'left'">
					<image v-if="item.contentType!=0" class="img" :src="item.headImg" mode="" @tap="linkToBusinessCard(item.fromId)"></image>
					<!-- contentType = 0 系统消息 -->
					<view class="system_content" v-if="item.contentType == 0">{{ item.content }}于 {{ getTime(item.createTime) }}</view>
					<!-- contentType = 1 文本 -->
					<view class="content" v-if="item.contentType == 1">{{ item.content }}</view>
					<!-- contentType = 2 语音 -->
					<view class="content contentType2" :class="[{ 'content-type-right': item.isItMe }]"
						v-if="item.contentType == 2" @tap="handleAudio(item)" hover-class="contentType2-hover-class"
						:style="{width:`${130+(item.contentDuration*2)}rpx`}">
						<view class="voice_icon" :class="[
								{ voice_icon_right: item.isItMe },
								{ voice_icon_left: !item.isItMe },
								{ voice_icon_right_an: item.anmitionPlay && item.isItMe },
								{ voice_icon_left_an: item.anmitionPlay && !item.isItMe }
							]"></view>
						<view class="">{{ item.contentDuration }}''</view>
					</view>
					<!-- contentType = 3 图片 -->
					<view class="content contentType3" v-if="item.contentType == 3" @tap="viewImg([item.temp])">
						<image :src="item.temp" class="img" mode="widthFix"></image>
					</view>
					<view v-if="item.contentType!=0" style="z-index: 0;color: #dddddd;margin: 0 8px 0 8px;font-size: 13px;">{{ getTime(item.createTime) }}</view>
				</view>
			</view>
		</view>

		<!-- 底部聊天输入框 -->
		<view class="input-box" :class="{ 'input-box-mpInputMargin': mpInputMargin }">
			<view class="input-box-flex">
				<!-- #ifndef H5 -->
				<image v-if="chatType === 'voice'" class="icon_img" :src="require('@/static/voice.png')"
					@click="switchChatType('keyboard')"></image>
				<image v-if="chatType === 'keyboard'" class="icon_img" :src="require('@/static/keyboard.png')"
					@click="switchChatType('voice')"></image>
				<!-- #endif -->
				<view class="input-box-flex-grow">
					<input v-if="chatType === 'voice'" type="text" class="content" id="input" v-model="formData.content"
						:hold-keyboard="true" :confirm-type="'send'" :confirm-hold="true"
						placeholder-style="color:#DDDDDD;" :cursor-spacing="10" @confirm="sendMsg(null)" />
					<view class="voice_title" v-if="chatType === 'keyboard'"
						:style="{ background: recording ? '#c7c6c6' : '#FFFFFF' }" @touchstart.stop.prevent="startVoice"
						@touchmove.stop.prevent="moveVoice" @touchend.stop="endVoice" @touchcancel.stop="cancelVoice">
						{{ voiceTitle }}
					</view>
				</view>

				<!-- 功能性按钮 -->
				<image class=" icon_btn_add" :src="require('@/static/add.png')" @tap="switchFun"></image>

				<!-- #ifdef H5 -->
				<button class="btn" type="primary" size="mini" @click="sendMsg(null)">发送</button>
				<!-- #endif -->
			</view>

			<view class="fun-box" :class="{'show-fun-box':showFunBtn}">
				<u-grid :col="4" hover-class="contentType2-hover-class" :border="false" @click="clickGrid">
					<u-grid-item v-for="(item, index) in funList" :index="index" :key="index" bg-color="#eaeaea">
						<u-icon :name="item.icon" :size="52"></u-icon>
						<view class="grid-text">{{ item.title }}</view>
					</u-grid-item>
				</u-grid>
			</view>

		</view>

		<!-- //语音动画 -->
		<view class="voice_an" v-if="recording">
			<view class="voice_an_icon">
				<view id="one" class="wave"></view>
				<view id="two" class="wave"></view>
				<view id="three" class="wave"></view>
				<view id="four" class="wave"></view>
				<view id="five" class="wave"></view>
				<view id="six" class="wave"></view>
				<view id="seven" class="wave"></view>
			</view>
			<view class="text">{{voiceIconText}}</view>
		</view>
	</view>
</template>

<script>
	import {
		getHost,
		concat_url,
		getTime,
	} from "requests/request"
	import {
		mapState
	} from 'vuex'
	export default {
		data() {
			return {
				toInfo: {},
				ToUsers: {
					"name": "群聊",
					"users": []
				},
				host: getHost(),
				that: this,
				formData: {
					content: '',
					limit: 15,
					index: 1
				},
				messageList: [],
				messageType: 1,
				loading: true, //标识是否正在获取数据
				imgHeight: '1000px',
				mpInputMargin: false, //适配微信小程序 底部输入框高度被顶起的问题
				chatType: "voice", // 图标类型 'voice'语音 'keyboard'键盘
				voiceTitle: '按住 说话',
				Recorder: uni.getRecorderManager(),
				Audio: uni.createInnerAudioContext(),
				recording: false, //标识是否正在录音
				isStopVoice: false, //加锁 防止点击过快引起的当录音正在准备(还没有开始录音)的时候,却调用了stop方法但并不能阻止录音的问题
				voiceInterval: null,
				voiceTime: 0, //总共录音时长
				canSend: true, //是否可以发送
				chatMember: [],
				PointY: 0, //坐标位置
				voiceIconText: "正在录音...",
				showFunBtn: false, //是否展示功能型按钮
				AudioExam: null, //正在播放音频的实例
				funList: [{
						icon: "photo-fill",
						title: "照片",
						uploadType: ["album"]
					},
					{
						icon: "camera-fill",
						title: "拍摄",
						uploadType: ["camera"]
					},
				],
			};
		},
		// uni.setNavigationBarTitle({
		//                 title: '我是新名字'//此处写页面的title
		//             });
		computed: {
			...mapState(['recsocket'])
		},
		watch: {
			"recsocket": function(val, old) {
				this.recMsg(val[val.length - 1])
			}
		},
		methods: {
			//拼接消息 处理滚动
			getTime(time){
				return getTime(time/1000000)
			},
			onNavigationBarButtonTap({ index }) {
				if (index == 0) {
					this.$u.route({
						url: 'pages/businessCard/businessCard',
						params: {
							userId:this.toInfo.toId,
							type: this.messageType,
						}
					});
				}
			},
			async joinData() {

				if (!this.loading) {
					//如果没有获取数据 即loading为false时，return 避免用户重复上拉触发加载
					return;
				}
				this.loading = false;
				const data = await this.getMessageData();
				if (data.length === 0) {

					this.loading = true;
					return
				}
				//获取节点信息
				const {
					index
				} = this.formData;
				const sel =
					`#msg-${index > 1 ? this.messageList[0].hasBeenSentId : data[data.length - 1].hasBeenSentId}`;
				this.messageList = [...data, ...this.messageList];
				//填充数据后，视图会自动滚动到最上面一层然后瞬间再跳回bindScroll的指定位置 ---体验不是很好，后期优化
				this.$nextTick(() => {
					this.bindScroll(sel);
					//如果还有数据
					if (this.formData.limit >= data.length) {
						this.formData.index++;
						setTimeout(() => {
							this.loading = true;
						}, 200);
					}
				});
			},
			recMsg(data) {
				data.isItMe = this._user_info.id.toString() == data.fromId
				let user = this.SearchInfo(data.fromId)
				data.headImg = concat_url(this.host, data.isItMe ? this._user_info.headImg : user.headImg)
				data.temp = this.getSrc(data)
				this.messageList.push(data);
				this.$nextTick(() => {
					// #ifdef MP-WEIXIN
					if (params.contentType == 1) {
						uni.pageScrollTo({
							scrollTop: 99999,
							duration: 0, //小程序如果有滚动效果 input的焦点也会随着页面滚动...
						});
					} else {
						setTimeout(() => {
							uni.pageScrollTo({
								scrollTop: 99999,
								duration: 0, //小程序如果有滚动效果 input的焦点也会随着页面滚动...
							});
						}, 150)
					}
					// #endif

					// #ifndef MP-WEIXIN
					uni.pageScrollTo({
						scrollTop: 99999,
						duration: 100
					});
					// #endif
				});

			},
			//处理滚动
			bindScroll(sel, duration = 0) {
				const query = uni.createSelectorQuery().in(this);
				query
					.select(sel)
					.boundingClientRect(data => {
						uni.pageScrollTo({
							scrollTop: data && data.top - 40,
							duration
						});
					})
					.exec();
			},
			//获取消息
			getMessageData() {
				var that = this;
				var url = "message/list?MessageType=" + that.messageType + "&Offset=" + that.messageList.length + "&ToId=" + that.toInfo.toId
				let getData = () => {
					let arr = [];
					that.request({
						url: url,
						method: "GET"
					}).then(res => {
						if (res.code === 200) {
							var msgs = res.data
							if (msgs === null || msgs.length == 0) {
								return arr
							}
							for (let i = 0; i < msgs.length; i++) {
								var isItMe = msgs[i].fromId == that._user_info.id
								let user = that.toInfo
								if (that.messageType == 2) {
									user = that.SearchInfo(msgs[i].fromId)
								}
								var obj = {
									hasBeenSentId: msgs[i].hasBeenSentId, //已发送过去消息的id
									content: msgs[i].content,
									fromId: msgs[i].fromId,
									headImg: isItMe ? that._user_info.headImg : concat_url(that.host,
										user.headImg), //用户头像
									toId: msgs[i].toId,
									isItMe, //true此条信息是我发送的 false别人发送的
									createTime: msgs[i].createAt,
									contentDuration: msgs[i].contentDuration,
									contentType: msgs[i].contentType, // 1文字文本 2语音
									pic: msgs[i].pic,
									anmitionPlay: false //标识音频是否在播放
								}
								obj.temp = that.getSrc(msgs[i])
								arr.unshift(obj);
							}
						}
					})

					return arr;
				};
				return new Promise((resolve, reject) => {
					const data = getData();
					setTimeout(() => {
						resolve(data);
					}, 500);
				});
			},
			//切换语音或者键盘方式
			switchChatType(type) {
				this.chatType = type;
				this.showFunBtn = false;
			},
			//切换功能性按钮
			switchFun() {
				this.chatType = 'keyboard'
				this.showFunBtn = !this.showFunBtn;
				uni.hideKeyboard()
			},
			//发送消息
			sendMsg(data) {
				console.log("点击send")
				// if (this.formData.content === ""  && data!=null && typeof data.content === 'undefined') {
				// 	return
				// }
				const params = data ? data : {
					hasBeenSentId: Date.now().toString(), //已发送过去消息的id
					content: this.formData.content,
					headImg: this._user_info.headImg, //用户头像
					fromId: this._user_info.id.toString(),
					messageType: this.messageType,
					toId: this.toInfo.toId.toString(),
					createTime: Date.now(),
					contentType: 1,
					temp: this.formData.content,
				};
				if (data) {
					if (data.contentType == 2) {
						//说明是发送语音
						params.content = data.content;
						params.contentType = data.contentType;
						params.contentDuration = data.contentDuration;
						params.anmitionPlay = false;
					} else if (data.contentType == 3) {
						//发送图片
						params.content = data.content;
						params.contentType = data.contentType;
					}
				} else if (!this.$u.trim(this.formData.content)) {
					//验证输入框书否为空字符传
					return;
				}
				this.formData.content = '';
				console.log("准备发送")
				this.$store.dispatch('WEBSOCKET_SEND', JSON.stringify(params))
				console.log("发送成功")
				params.isItMe = this._user_info.id.toString() == params.fromId
				this.messageList.push(params);

				this.$nextTick(() => {
					// #ifdef MP-WEIXIN
					if (params.contentType == 1) {
						uni.pageScrollTo({
							scrollTop: 99999,
							duration: 0, //小程序如果有滚动效果 input的焦点也会随着页面滚动...
						});
					} else {
						setTimeout(() => {
							uni.pageScrollTo({
								scrollTop: 99999,
								duration: 0, //小程序如果有滚动效果 input的焦点也会随着页面滚动...
							});
						}, 150)
					}
					// #endif

					// #ifndef MP-WEIXIN
					uni.pageScrollTo({
						scrollTop: 99999,
						duration: 100
					});
					// #endif

					if (this.showFunBtn) {
						this.showFunBtn = false;
					}

					// #ifdef MP-WEIXIN 
					if (params.contentType == 1) {
						this.mpInputMargin = true;
					}
					// #endif
					//h5浏览器并没有很好的办法控制键盘一直处于唤起状态 而且会有样式性的问题
					// #ifdef H5
					uni.hideKeyboard();
					// #endif
				});
			},
			//用户触摸屏幕的时候隐藏键盘
			touchstart() {
				uni.hideKeyboard();
			},
			// userid 用户id
			linkToBusinessCard(userId) {
				this.$u.route({
					url: 'pages/businessCard/businessCard',
					params: {
						userId:userId,
						type: 1,
					}
				});
			},
			//准备开始录音
			startVoice(e) {
				if (!this.Audio.paused) {
					//如果音频正在播放 先暂停。
					this.stopAudio(this.AudioExam)
				}
				this.recording = true;
				this.isStopVoice = false;
				this.canSend = true;
				this.voiceIconText = "正在录音..."
				this.PointY = e.touches[0].clientY;
				this.Recorder.start({
					format: 'mp3'
				});
			},
			//录音已经开始
			beginVoice() {
				if (this.isStopVoice) {
					this.Recorder.stop();
					return;
				}
				this.voiceTitle = '松开 结束'
				this.voiceInterval = setInterval(() => {
					this.voiceTime++;
				}, 1000)
			},
			//move 正在录音中
			moveVoice(e) {
				const PointY = e.touches[0].clientY
				const slideY = this.PointY - PointY;
				if (slideY > uni.upx2px(120)) {
					this.canSend = false;
					this.voiceIconText = '松开手指 取消发送 '
				} else if (slideY > uni.upx2px(60)) {
					this.canSend = true;
					this.voiceIconText = '手指上滑 取消发送 '
				} else {
					this.voiceIconText = '正在录音... '
				}
			},
			//结束录音
			endVoice() {
				this.isStopVoice = true; //加锁 确保已经结束录音并不会录制
				this.Recorder.stop();
				this.voiceTitle = '按住 说话'
			},
			//录音被打断
			cancelVoice(e) {
				this.voiceTime = 0;
				this.voiceTitle = '按住 说话';
				this.canSend = false;
				this.Recorder.stop();
			},
			//处理录音文件
			handleRecorder({
				tempFilePath,
				duration
			}) {
				let contentDuration;
				// #ifdef MP-WEIXIN
				this.voiceTime = 0;
				if (duration < 600) {
					this.voiceIconText = "说话时间过短";
					setTimeout(() => {
						this.recording = false;
					}, 200)
					return;
				}
				contentDuration = duration / 1000;
				// #endif

				// #ifdef APP-PLUS
				contentDuration = this.voiceTime + 1;
				this.voiceTime = 0;
				if (contentDuration <= 0) {
					this.voiceIconText = "说话时间过短";
					setTimeout(() => {
						this.recording = false;
					}, 200)
					return;
				};
				// #endif
				this.recording = false;
				var that = this
				uni.uploadFile({
					url: "http://" + that.host + 'file/upload',
					filePath: tempFilePath,
					name: "file",
					fromData: {
						'name': 'temp.mp3'
					},
					header: {
						"Content-Type": "multipart/form-data"
					},
					success: uploadFileRes => {
						if (uploadFileRes.statusCode === 200) {
							var data = JSON.parse(uploadFileRes.data).data;
							const params = {
								contentType: 2,
								content: data,
								contentDuration: Math.ceil(contentDuration),
								hasBeenSentId: Date.now().toString(), //已发送过去消息的id
								headImg: this._user_info.headImg,
								fromId: this._user_info.id.toString(),
								messageType: this.messageType,
								toId: this.toInfo.toId.toString(),
								createTime: Date.now(),
								temp: tempFilePath,
							};
							this.canSend && this.sendMsg(params);
						} else {
							uni.showToast({
								title: '发送失败',
								icon: 'none'
							});
						}
					},
					fail: res => {

						uni.showToast({
							title: '上传失败',
							icon: 'none'
						});
						uni.hideLoading();
					}
				});

			},
			//控制播放还是暂停音频文件
			handleAudio(item) {
				this.AudioExam = item;
				this.Audio.paused ? this.playAudio(item) : this.stopAudio(item);
			},
			//播放音频
			playAudio(item) {
				this.Audio.src = item.temp;
				this.Audio.hasBeenSentId = item.hasBeenSentId;
				this.Audio.play();
				item.anmitionPlay = true;
			},
			getSrc(item) {
				if (item.contentType == 1) {
					return
				}
				if (typeof item.temp === 'undefined') {
					return concat_url(this.host, item.content);
				} else {
					return item.content;
				}
			},
			//停止音频
			stopAudio(item) {
				item.anmitionPlay = false;
				this.Audio.src = '';
				this.Audio.stop();
			},
			//关闭动画
			closeAnmition() {
				const hasBeenSentId = this.Audio.hasBeenSentId;
				const item = this.messageList.find(it => it.hasBeenSentId == hasBeenSentId);
				item.anmitionPlay = false;
			},
			//点击宫格时触发
			clickGrid(index) {
				if (index == 0) {
					this.chooseImage(['album'])
				} else if (index == 1) {
					this.chooseImage(['camera'])
				}
			},
			SearchInfo(id) {
				if (this.messageType == 1) {
					return this.toInfo
				}
				for (var i = 0; i < this.chatMember.length; i++) {
					if (this.chatMember[i].id === id) {
						return this.chatMember[i]
					}
				}
				return this.toInfo
			},
			//发送图片
			chooseImage(sourceType) {
				var that = this
				uni.chooseImage({
					sourceType,
					sizeType: ['compressed'],
					success: res => {
						this.showFunBtn = false;
						for (let i = 0; i < res.tempFilePaths.length; i++) {
							uni.uploadFile({
								url: "http://" + that.host + 'file/upload',
								filePath: res.tempFilePaths[i],
								name: "file",
								fromData: {
									'name': 'temp.jpg'
								},
								header: {
									"Content-Type": "multipart/form-data"
								},
								success: uploadFileRes => {
									if (uploadFileRes.statusCode === 200) {
										var data = JSON.parse(uploadFileRes.data).data;
										const params = {
											contentType: 3,
											content: data,
											hasBeenSentId: Date.now().toString(), //已发送过去消息的id
											headImg: this._user_info.headImg,
											fromId: this._user_info.id.toString(),
											messageType: this.messageType,
											toId: this.toInfo.toId.toString(),
											createTime: Date.now(),
											temp: res.tempFilePaths[i],
										};
										this.sendMsg(params)
									} else {
										uni.showToast({
											title: '发送失败',
											icon: 'none'
										});
									}
								},
								fail: res => {

									uni.showToast({
										title: '上传失败',
										icon: 'none'
									});
									uni.hideLoading();
								}
							});
						}
					}
				})
			},
			//查看大图
			viewImg(imgList) {
				uni.previewImage({
					urls: imgList,
					// #ifndef MP-WEIXIN
					indicator: 'number'
					// #endif
				});
			},
		},
		onPageScroll(e) {
			if (e.scrollTop < 50) {
				this.loading = true
				this.joinData();
			}
		},
		onNavigationBarButtonTap({
			index
		}) {
			if (index == 0) {
				//用户详情 设置
			} else if (index == 1) {
				//返回按钮
				this.$u.route({
					url: 'pages/home/home'
				});
			}
		},

		//返回按钮事件
		onBackPress(e) {
			//以下内容对h5不生效
			//--所以如果用浏览器自带的返回按钮进行返回的时候页面不会重定向 正在寻找合适的解决方案
			this.$u.route({
				type: "switchTab",
				url: 'pages/home/home'
			});
			return true;
		},
		onLoad(info) {
			this.toInfo = {
				toName: info.toName,
				headImg: info.headImg,
				toId: info.relationId,
			};
			this.messageType = Number(info.messageType)
			//录音开始事件
			this.Recorder.onStart(e => {

				this.beginVoice();
			});
			//录音结束事件
			this.Recorder.onStop(res => {
				clearInterval(this.voiceInterval);
				this.handleRecorder(res);
			});

			//音频停止事件
			this.Audio.onStop(e => {
				this.closeAnmition();
			});

			//音频播放结束事件
			this.Audio.onEnded(e => {
				this.closeAnmition();
			});
		},


		onReady() {
			// this.$store.dispatch('WEBSOCKET_INIT', 'ws://localhost:8081/socket')
			// 发送websocket数据：
			// this.$store.dispatch('WEBSOCKET_SEND', '发送的数据')
			// 关闭websocket连接：
			// this.$store.dispatch('CLOSE_SOCKET')

			//自定义返回按钮 因为原生的返回按钮不可阻止默认事件
			// #ifdef H5
			const icon = document.getElementsByClassName('uni-page-head-btn')[0];
			icon.style.display = 'none';
			// #endif

			uni.setNavigationBarTitle({
				title: this.toInfo.toName
			});
			var url = "user/fromMessage?MessageType=" + this.messageType + "&ToId=" + this.toInfo.toId
			var that = this;
			this.request({
				url: url,
				method: "GET"
			}).then(res => {
				if (res.code === 200) {
					var data = res.data
					that.chatMember = data
					that.joinData();
				}
			})

			uni.getSystemInfo({
				success: res => {
					this.imgHeight = res.windowHeight + 'px';
				}
			});

			uni.onKeyboardHeightChange(res => {
				if (res.height == 0) {
					// #ifdef MP-WEIXIN
					this.mpInputMargin = false;
					// #endif
				} else {
					this.showFunBtn = false;
				}
			});
		}
	};
</script>

<style lang="scss" scoped>
	@import './index.scss'
	
</style>
