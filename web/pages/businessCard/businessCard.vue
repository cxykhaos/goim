<template>
	<view>
		<u-action-sheet :list="list" @click="clickList"  v-model="show"></u-action-sheet>
		<view v-if="type!==1" class="groupMember">
			<view class="memberInfo" v-for="(item, index) in groupInfo.groupMembers" >
				<image :src="item.headImg" @tap="now=item;show=true;" class="img"></image>
				<view style="text-align: center;">{{item.username}}</view>
			</view>
		</view>
		
		<view class="userinfo" v-if="type!==1">
			<image :src="groupInfo.headImg" @tap="previewImg(groupInfo.headImg)" class="img"></image>
				<view class="userinfo-desc">
					<view class="userinfo-desc-name">{{groupInfo.groupName}}</view>
					<view class="userinfo-desc-gray">群昵称：{{groupInfo.groupName}}</view>
					<view class="userinfo-desc-gray">微信号：{{groupInfo.id}}</view>
				</view>
		</view>
		<view class="userinfo" v-if="type===1">
			<image :src="userInfo.headImg" @tap="previewImg(userInfo.headImg)" class="img"></image>
			<view class="userinfo-desc">
				<view class="userinfo-desc-name">{{userInfo.username}}</view>
				<view class="userinfo-desc-gray">昵称：{{userInfo.username}}</view>
				<view class="userinfo-desc-gray">微信号：{{userInfo.id}}</view>
				<view class="userinfo-desc-gray">地区：{{userInfo.address}}</view>
			</view>
		</view>
		<view  class="perch"></view>
		<u-cell-group  v-if="type===1">
			<u-cell-item title="朋友圈" label="模拟数据暂不支持查看好友朋友圈" hover-class="none" :title-style="{ marginLeft: '10rpx' }"></u-cell-item>
			<u-cell-item title="更多信息" :title-style="{ marginLeft: '10rpx' }" @click="linkToMoreInfoMation"></u-cell-item>
		</u-cell-group>
		<view class="" v-if="!isItMe">
			<view  class="perch"></view>
			<u-cell-group >
				<u-cell-item title="发消息" :arrow="false" :center="true" :title-style="{ marginLeft: '10rpx' }" @click="linkToChat">
					<u-icon slot="icon" name="chat" color="#36648B" size="32"></u-icon>
				</u-cell-item>
			</u-cell-group>
		</view>
	</view>
</template>

<script>
	import {
		getHost,
		concat_url,
		removeById,
		getTime,
	} from "requests/request"
	export default {
		data() {
			return {
				userInfo:{},
				now:{},
				isItMe:false,
				selectType : "linkto",
				type:1,
				groupInfo:{},
				list: [ {
						text: '查看信息'
				}],
				show: false
			}
		},
		methods: {
			linkToChat(){
				let params = {}
				if (this.type===1){
					params = {
						relationId:this.userInfo.id,headImg:this.userInfo.headImg,toName:this.userInfo.username,messageType: 1
					}
				}else{
					params = {
						relationId:this.groupInfo.id,headImg:this.groupInfo.headImg,toName:this.groupInfo.groupName,messageType: 2
					}
				}
				this.$u.route({
					url:"pages/chat/chat",
					params: params
				})
			}, 
			linkToBusinessCard(userId) {
				this.$u.route({
					url: 'pages/businessCard/businessCard',
					params: {
						userId:userId,
						type: 1,
					}
				});
			},
			previewImg(urls){
				uni.previewImage({urls:[urls]})
			},
			clickList(index) {
				if (index === 0){
					this.linkToBusinessCard(this.now.id)
				}else if (index ===1){
					removeById(this.groupInfo.groupMembers,this.now.id)
				}
			},
			linkToMoreInfoMation(){
				this.$u.route({
					url:"pages/moreInforMation/moreInforMation",
					params:{ signature:this.userInfo.signature }
				})
			}
		},
		onLoad({ userId,type }) {
			var that = this
			that.type = type === "1"?1:0
			if( type === "1" && userId === this._user_info.id){
				this.userInfo = this._user_info;
				this.isItMe = true;
				return;
			}
			var url = "user/otherInfo?Type=" + type + "&ToId=" + userId
			that.request({
				url: url,
				method: "GET"
			}).then(res => {
				if (res.code === 200) {
					let data =res.data
					if (type === "1"){
						that.userInfo = data
						that.userInfo.headImg = concat_url(getHost(),data.headImg)
					}else{
						data.headImg = concat_url(getHost(),data.headImg)
						for (let i=0;i<data.groupMembers.length;i++){
							data.groupMembers[i].headImg = concat_url(getHost(),data.groupMembers[i].headImg)
							that.groupInfo = data
						}
						if (data.ownId === that._user_info.id){
							that.list.push({
									text: '删除成员',
									color: 'red',
								})
						}
					}
					
				}
					
			})
		}
	}
</script>

<style scoped lang="scss">
	.perch{
		height: 10rpx;
	}
.groupMember{
	display: flex;  
	flex-wrap: wrap;
	.memberInfo{
		margin: 10rpx;
	}
	.img{
		display: block;
		width: 88rpx;
		height: 88rpx;
		border-radius: 10rpx;
		margin-bottom: 10rpx;
	}
	padding: 20rpx 20rpx 40rpx 20rpx;
	background-color: #FFFFFF;
	
}
.userinfo{
	display: flex;
	justify-content: flex-start;
	align-items: flex-start;
	padding: 20rpx 30rpx 40rpx 40rpx;
	background-color: #FFFFFF;
	.img{
		display: block;
		width: 110rpx;
		height: 110rpx;
		border-radius: 10rpx;
	}
	&-desc{
		padding-left: 30rpx;
		&-name{
			font-weight: bold;
			font-size: 36rpx;
			transform: translateY(-10rpx);
		}
		&-gray{
			color: $u-tips-color;
			font-size: 29rpx;
		}
	}
}
</style>
