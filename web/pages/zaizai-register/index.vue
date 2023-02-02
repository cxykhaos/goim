<template>
	<view class="zai-box">
		<image src="../../static/khaos.png" mode='aspectFit' class="zai-logo"></image>
<!-- 		<view class="zai-title">LOGO区域</view> -->
		<view class="zai-form">
			<input class="zai-input" v-model="userId" placeholder="请输入用户名" />
			<input class="zai-input" v-model="userPwd" password placeholder="请输入密码"/>
			<input class="zai-input" v-model="doublePwd" password placeholder="请再输入一次密码"/>
			<button class="zai-btn" @click="register">立即注册</button>
			<navigator url="../zaizai-login/index" open-type='navigateBack' hover-class="none" class="zai-label">已有账号，点此去登录.</navigator>
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
				userPwd:"",
				doublePwd:"",
			}
		},
		methods:{
			CheckNotVaild(){
				if (this.userId===""){
					this.$u.toast('账号不能为空');
					return true
				} 
				if (this.userPwd==="" ){
					this.$u.toast('密码不能为空');
					return true
				}
				
				if (this.userPwd !== this.doublePwd){
					this.$u.toast('两次密码不一致');
					return true
				}
				
				return false
			},
			register(){
				var that = this;
				if (that.CheckNotVaild()){
					return
				}
				var url = "http://"+getHost()+"user/register"
				
				uni.request({
					url: url,
					method: 'POST',
					data:{userId:that.userId},
					success(res) {
						let data = res.data
						if (data.code == 200) {
							that.$u.toast('注册成功');
							that.$u.route({
								type: 'reLaunch',
								url: 'pages/zaizai-login/index',
								params:{ userId:that.userId,userPwd:that.userPwd }
								
							});
						}else{
							that.$u.toast(res.data.msg);
						}
					},
					fail(res) {
						console.log(res)
					}
				})
			},
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
		margin-top: 60upx;
	}
	.zai-btn:after{
		border: 0;
	}
	/*按钮点击效果*/
	.zai-btn.button-hover{
		transform: translate(1upx, 1upx);
	}
</style>
