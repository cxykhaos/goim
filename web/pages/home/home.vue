<template>
	<view class="content">
		<!-- #ifdef MP-WEIXIN -->
		<u-navbar :is-back="false" title="微信" :background="{ background: '#f8f8f8'  }" title-color="#404133" :border-bottom="false" z-index="1001">
			<view class="slot-wrap" slot="right">
				<u-icon name="plus-circle" size="36" @click="showSelect"></u-icon>
			</view>		
		</u-navbar>
		<!-- #endif -->
		
		<selectInput :list="selectList" :list-key="'name'" :show.sync="selectShow" @on-select="checkSelect" @close="closeSelect" />
		<searchInput />
		<u-swipe-action :show="item.show" :index="index" v-for="(item, index) in friendslist" btn-width="160" :key="item.id" @click="click" @open="open" :options="options">
			<view class="item" :class="item.isTop ? 'bg_view' : ''" hover-class="message-hover-class" @tap="linkTo(item)">
				<image mode="aspectFill" :src="item.images" />
				<!-- 此层wrap在此为必写的，否则可能会出现标题定位错误 -->
				<view class="right u-border-bottom title-wrap">
					<view class="right_top">
						<view class="right_top_name u-line-1">{{ item.name }}</view>
						<view class="right_top_time ">{{ item.updateTime }}</view>
					</view>
					<view class="right_btm ">
						<view class="u-line-1">{{ item.NewMsg }}</view>
						<view class="red-sign" v-if="item.unread!=0">{{ item.unread }}</view>
					</view>
				</view>
			</view>
		</u-swipe-action>
		
	</view>
</template>

<script>
import searchInput from '@/components/searchInput/index.vue';
import selectInput from '@/components/selectInput/selectInput.vue';
import { mapState } from 'vuex';
import { concat_url,getHost,getTime } from "requests/request"
export default {
	components: { searchInput, selectInput },
	data() {
		return {
			host: getHost(),
			friendslist: [],
			selectShow: false,
			options: [
				{
					text: '删除',
					style: {
						backgroundColor: '#dd524d',
						fontSize: '24rpx'
					}
				}
			],
			selectList: [{ id: '1', name: '添加朋友', icon: 'man-add-fill' }, { id: '2', name: '扫一扫', icon: 'scan' }, { id: '3', name: '收付款', icon: 'fingerprint' }]
		};
	},
	computed: {
			...mapState(['recsocket'])
	},
	watch: {
		"recsocket": function(val, old) {
			this.recNewMsg(val[val.length-1])
		}
	},
	mounted(){
		this.$store.dispatch('WEBSOCKET_INIT', 'ws://'+ getHost() +'socket')
		let routes = getCurrentPages()
			var that = this
			this.request({
				url: "user/info",
				method: "GET"
			}).then(res => {
				if (res.code === 200) { 
					let data = res.data
					data.headImg= concat_url(getHost(),data.headImg)
					data.chatBgImg= concat_url(getHost(),data.chatBgImg)
					that.$store.commit("CHANGE_INFO",data)
				} 
			})
					
	},
	onShow(){
		var that = this;
		var arr = []
		var url = "http://" + that.host+ "relation/list"
		this.request({
			url: "relation/list",
			method: "GET"
		}).then(res => {
			if (res.code === 200) { 
				var data = res.data
				var arr = []
				if (data === null || data === [] || data.length == 0) {
					return arr
				}
				for (let i = 0; i < data.length; i++) {
					data[i].relationId = data[i].relationId
					data[i].id = i
					data[i].name = data[i].name
					data[i].images = concat_url(getHost(),data[i].headImg);
					let r = that.getPreviewContent(data[i].newMsg)
					data[i].updateTime=r[0]
					data[i].NewMsg=r[1]
					data[i].unread=0
					data[i].show= false
					arr.push(data[i])
				}
				arr.sort(function(a,b){return b.newMsg.createAt - a.newMsg.createAt;});
				that.friendslist = arr
				}})
		
	},
	methods: {
		recNewMsg(data) {
			for (let i=0;i<this.friendslist.length;i++){
				let stand = data.fromId
				if (data.messageType==2){
					stand = data.toId
				}
				if (this.friendslist[i].relationId === stand && this.friendslist[i].messageType===data.messageType){
					let r = this.getPreviewContent(data)
					this.friendslist[i].updateTime = r[0]
					this.friendslist[i].NewMsg= r[1]
					this.friendslist[i].unread=this.friendslist[i].unread+1
					let temp = this.friendslist[i]
					this.friendslist.splice(i, 1);
					this.friendslist.unshift(temp)
					break;
				}
				
			}
		},
		getPreviewContent(data){
			let t = getTime(data.createAt/1000000)
			switch(data.contentType){
				case 0: case 1:
					return [t,data.content]
				case 2:
					return [t,"[语音]"]
				case 3:
					return [t,"[图片]"]
				default:
					return [""," ㅤ"]
			}	
		},
		showSelect(){
			this.selectShow = !this.selectShow;
		},
		//action 点击事件
		click(index, index1) {
			if (index1 == 0) {
				this.friendslist.splice(index, 1);
			} 
		},
		//action 打开事件
		open(index) {
			this.friendslist[index].show = true;
			this.friendslist.map((val, idx) => {
				if (index != idx) this.friendslist[idx].show = false;
			});
		},
		//点击导航栏自定义按钮
		onNavigationBarButtonTap({ index }) {
			if (index == 0) {
				this.showSelect()
			}
		},
		//跳转
		linkTo(item) {
			this.$u.route({
				url: 'pages/chat/chat',
				params: { relationId:item.relationId,headImg:item.headImg,toName:item.name,messageType: item.messageType}
			});
		},
		//关闭弹窗
		closeSelect(){
			//小程序兼容
			this.selectShow = false;
		},
		//扫码
		checkSelect(index) {
			var that = this
			if (index == 1) {
				uni.vibrateLong();
				//扫码
				uni.scanCode({
					success: function(res) {
						// console.log('条码内容：' + res.result);
						this.$u.toast('条码类型:', res.scanType);
					}
				});
			}else if (index == 0){
				this.$u.route({
					url:"pages/search/search",
					params: { type:"add"}
				})
				
			}
		},
	},
};
</script>

<style lang="scss" scoped>
.content {
	.item {
		width: 750rpx;
		display: flex;
		align-items: center;
		// padding: 20rpx;
		image {
			width: 76rpx;
			height: 76rpx;
			margin: 20rpx;
			border-radius: 12rpx;
			flex: 0 0 76rpx;
		}
		.right {
			overflow: hidden;
			flex: 1 0;
			padding: 20rpx 20rpx 20rpx 0;
			&_top {
				display: flex;
				justify-content: space-between;
				&_name {
					font-size: 28rpx;
					color: $u-main-color;
					flex: 0 0 400rpx;
					overflow: hidden;
				}
				&_time {
					font-size: 22rpx;
					color: $u-light-color;
				}
			}
			&_btm {
				display: flex;
				justify-content: space-between;
				align-items: center;
				font-size: 22rpx;
				color: $u-tips-color;
				padding-top: 10rpx;
			}
		}
	}
	.bg_view {
		background-color: #fafafa;
	}
	.slot-wrap {
		display: flex;
		align-items: center;
		padding: 0 30rpx; 
	}
	.red-sign{
		    background-color: red;
		    color: white;
		    padding: 1px 8px 0 8px;
		    border-radius: 16px;
	}
}
</style>
