<template>
	<view class="content">
		<!-- #ifdef MP-WEIXIN -->
		<u-navbar :is-back="false" title=" " :background="{ background: '#ffffff'  }" :border-bottom="false"></u-navbar>
		<!-- #endif -->
		<!-- #ifndef MP-WEIXIN -->
		<view class="status_bar"></view>
		<!-- #endif -->
		<view class="header">
			<view class="header_left">
				<image mode="aspectFill" :src="_user_info.headImg" @tap="uploadImg" />
				<view class="header_left_content">
					<view class="header_left_content_name">{{_user_info.username}}</view>
					<view class="header_left_content_number">微信号: {{_user_info.id}}</view>
				</view>
			</view>
			<view class="header_right">
				<u-icon name="arrow-right"></u-icon>
			</view>
		</view>
		<view style="height: 10rpx;"></view>
		<u-cell-group>
			<u-cell-item title="支付" :title-style="{ marginLeft: '10rpx' }">
				<u-icon slot="icon" name="fingerprint" color="#4cd964" size="32"></u-icon>
			</u-cell-item>
		</u-cell-group>
		<view style="height: 10rpx;"></view>
		<u-cell-group>
			<u-cell-item v-for="(item, index) in groupList" :key="index" :title="item.title"
				:title-style="{ marginLeft: '10rpx' }">
				<u-icon slot="icon" :name="item.icon" :color="item.color" size="32"></u-icon>
			</u-cell-item>
		</u-cell-group>
		<view style="height: 10rpx;"></view>
		<u-cell-group>
			<u-cell-item title="退出账号" :title-style="{ marginLeft: '10rpx' }" @click="quit">
				<u-icon slot="icon" name="setting" color="#409eff" size="32"></u-icon>
			</u-cell-item>
		</u-cell-group>
	</view>
</template>

<script>
import { getHost,getTime } from "requests/request"
	export default {
		data() {
			return {
			host: getHost(),
				groupList: [{
						title: '收藏',
						color: '#409eff',
						icon: 'star'
					},
					{
						title: '相册',
						color: '#409eff',
						icon: 'photo'
					},
					{
						title: '卡包',
						color: '#409eff',
						icon: 'coupon'
					},
					{
						title: '表情',
						color: '#ff9900',
						icon: 'heart'
					}
				],
			};
		},
		onShow: function(){
			console.log(this._user_info)
		},
		methods: {
			quit(){
				console.log(1)
				uni.removeStorageSync('token')
				this.$u.route({
					type: 'reLaunch',
					url: 'pages/zaizai-login/index'
				});
				console.log(2)
			},
			uploadImg() {
				var that = this
				uni.chooseImage({
					count: 1,
					sizeType: ['compressed'],
					success: res => {
						console.log("http://" + that.host + 'file/upload')
						uni.uploadFile({
							url: "http://" + that.host + 'file/upload',
							formData: {
											'type': 'avatar'
										},
							header:{'token':uni.getStorageSync('token') || ''},
							filePath: res.tempFilePaths[0],
							name: "file",
							success: uploadFileRes => {
								if (uploadFileRes.statusCode === 200) {
									var data = JSON.parse(uploadFileRes.data).data;
									that._user_info.headImg = res.tempFilePaths[0]

								} else {
									uni.showToast({
										title: '头像更换失败',
										icon: 'none'
									});
								}
							},
							fail: res => {
								console.log()
								uni.showToast({
									title: '上传失败\n'+res,
									icon: 'none'
								});
								uni.hideLoading();
							}
						});
					}
				})
			}
		}
	};
</script>

<style lang="scss" scoped>
	.header {
		display: flex;
		justify-content: space-between;
		align-items: center;
		padding: 60rpx 38rpx;
		background-color: #ffffff;

		&_left {
			display: flex;
			align-items: center;

			&_content {
				padding-left: 20rpx;

				&_name {
					font-weight: bold;
				}

				&_number {
					color: #969799;
					font-size: 26rpx;
				}
			}
		}

		&_right {
			font-size: 28rpx;
			color: #969799;
		}

		image {
			width: 100rpx;
			height: 100rpx;
			border-radius: 6rpx;
		}
	}

	.status_bar {
		height: var(--status-bar-height);
		width: 100%;
	}
</style>
