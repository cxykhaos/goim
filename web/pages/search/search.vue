<template>
	<view class="">
		<!-- #ifdef MP-WEIXIN -->
		<u-navbar title=" " :background="{ background: '#f8f8f8' }">
			<view class="slot-wrap">
				<u-search
					:focus="true"
					@custom="clickCancel"
					placeholder="搜索..."
					shape="square"
					:action-style="{ color: '#007aff' }"
					action-text="取消"
					:bg-color="'#ffffff'"
					@clear="cleanSearch"
					v-model="search_word"
				></u-search>
			</view>
		</u-navbar>
		<!-- #endif -->
		<!-- #ifndef MP-WEIXIN -->
		<view class="status_bar"></view>
		<view class="content_search">
			<u-search
				:focus="true"
				@custom="clickCancel"
				placeholder="h5能否自动获取焦点取决于浏览器的实现"
				shape="square"
				:action-style="{ color: '#007aff' }"
				@clear="cleanSearch"
				action-text="取消"
				:bg-color="'#ffffff'"
				v-model="search_word"
			></u-search>
		</view>
		<!-- #endif -->
		<view :index="index" v-for="(item, index) in searchlist" btn-width="160" :key="item.id">
			<view class="item" :class="item.isTop ? 'bg_view' : ''" hover-class="message-hover-class" @tap="">
				<image mode="aspectFill" :src="item.images" />
				<!-- 此层wrap在此为必写的，否则可能会出现标题定位错误 -->
				<view class="right u-border-bottom title-wrap">
					<view class="right_top">
						<view class="right_top_name u-line-1">{{ item.name }}</view>
					
						<u-button type="success" @click="AddRelation(item.messageType,item.relationId)">添加</u-button>
					</view>
				</view>
			</view>
		</view>
	</view>
</template>

<script>
import { concat_url,getHost,getTime } from "requests/request"
export default {
	data() {
		return {
			search_word: '',
			search_type: 'search',
			searchlist:[],
		};
	},
	watch: {
		"search_word": function(val, old) {
			this.search(val)
			
		}
	},
	methods: {
		clickCancel() {
			this.$u.route({
				type: 'navigateBack'
			});
		},
		AddRelation(messageType,id){
			this.request({
				url: "relation/add",
				method: "POST",
				data: { messageType:messageType,relationId:id },
			}).then(res => {
				if (res.code === 200) { 
					uni.showToast({
						icon:'success',
						title:'添加成功'
					})
				}else{
					uni.showToast({
						icon:'error',
						title:'该关系已添加'
					})
				}
			})
		},
		cleanSearch(){
			this.searchlist=[]
		},
		search(val){
			var that = this
			this.request({
				url: "search/list?searchWord="+val,
				method: "GET",
			}).then(res => {
				if (res.code === 200) { 
					let data = res.data
					var arr = []
					if (data === null || data === [] || data.length == 0) {
						that.searchlist=arr
						return
					}
					for (let i = 0; i < data.length; i++) {
						data[i].relationId = data[i].relationId
						data[i].id = i
						data[i].name = data[i].name
						data[i].images =concat_url(getHost(),data[i].headImg)
						arr.push(data[i])
					}
					
					console.log(arr)
					that.searchlist=arr
				}
			})
			
		}
	},
	onShow() {},
	onLoad(info){
		this.search_type = info.type
	}
};
</script>

<style lang="scss" scoped>
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

.status_bar {
	height: var(--status-bar-height);
	width: 100%;
}
.content_search {
	padding: 16rpx;
	background-color: #f8f8f8;
}
</style>
