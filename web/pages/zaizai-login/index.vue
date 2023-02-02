<template>
	<view class="zai-box">
		<image src="../../static/khaos.png" mode='aspectFit' class="zai-logo"></image>
	<!-- 	<view class="zai-title">LOGO区域</view> -->
		<view class="zai-form">
			<input class="zai-input"  v-model="userId" placeholder="请输入用户名" />
			<input class="zai-input"  password placeholder="请输入密码"/>
			<view class="zai-label">忘记密码？</view>
			<button class="zai-btn" @click="login">立即登录</button>
			<navigator url="../zaizai-register/index" hover-class="none" class="zai-label">还没有账号？点此注册.</navigator>
		</view>
	</view>
</template>

<script>
	import { getHost,concat_url } from "requests/request"
	import store from 'store/index.js'
	export default {
		data(){
			return {
				userId:"",	
			}
		},
		mounted:function(){
			let token = uni.getStorageSync("token")
			if (token!==""){
				this.jumpToHome()
			}
		},
		onLaunch: function() {
		},
		onLoad(info){
			console.log(info)
		},
		onShow: function() {
			this.$store.dispatch('CLOSE_SOCKET')
			let routes = getCurrentPages()
			console.log(routes[routes.length - 1].route)
		},
		onHide: function() {
			
		},
		beforeDestroy: function(){
				this.$store.dispatch('CLOSE_SOCKET')
		},
		methods:{
			login(){
				var url = "http://"+getHost()+"user/login"
				console.log(url)
				var that = this
				uni.request({
					url: url,
					method: 'POST',
					data:{userId:that.userId},
					success(res) {
						console.log(res)
						if (res.statusCode == 200) {
							console.log(res)
							uni.setStorageSync('token',res.data.token)
							uni.setStorageSync('userId',res.data.userId)
							that.jumpToHome()
						}
					},
					fail(res) {
						console.log(res)
					}
				})
			},
			jumpToHome(){
			  this.$u.route({
			  	type: 'switchTab',
			  	url: 'pages/home/home'
			  });
			}  
		}
	}
	
</script>

<style>
	.zai-box{
		padding: 0 100upx;
		position: relative;
	}
	.zai-logo{
		width: 100%;
		width: 100%;
		height: 310upx;
	}
	.zai-title{
		position: absolute;
		top: 0;
		line-height: 360upx;
		font-size: 68upx;
		color: #fff;
		text-align: center;
		width: 100%;
		margin-left: -100upx;
	}
	.zai-form{
		margin-top: 300upx;
	}
	.zai-input{
		background: #e2f5fc;
		margin-top: 30upx;
		border-radius: 100upx;
		padding: 20upx 40upx;
		font-size: 36upx;
	}
	.input-placeholder, .zai-input{
		color: #94afce;
	}
	.zai-label{
		padding: 60upx 0;
		text-align: center;
		font-size: 30upx;
		color: #a7b6d0;
	}
	.zai-btn{
		background: #ff65a3;
		color: #fff;
		border: 0;
		border-radius: 100upx;
		font-size: 36upx;
	}
	.zai-btn:after{
		border: 0;
	}
	/*按钮点击效果*/
	.zai-btn.button-hover{
		transform: translate(1upx, 1upx);
	}
</style>
