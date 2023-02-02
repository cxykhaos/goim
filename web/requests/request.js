const BASE_URL = "http://"+getHost() 

export const request = (options => {
	return new Promise((resolve, reject) => {
		uni.request({
			url: BASE_URL + options.url,
			method: options.method || 'GET',
			header: {
				token: uni.getStorageSync('token') || ''
			},
			data: options.data || {},
			success: (res) => {
				const data = res.data
				if (res.statusCode >= 401 ) {
					
					uni.removeStorageSync('token')
				console.log(res.statusCode)
					uni.reLaunch({ 
								url: '/pages/zaizai-login/index',
								animationType: 'pop-in',
								animationDuration: 300
							})
					return
				} else if (res.statusCode !== 200) {
					uni.showToast({
						icon:'error',
						title:'操作错误'
					})
				}
				resolve(data)
			},
			fail:(error)=>{
				console.log(error)
				uni.showToast({
					icon:'error',
					title:'系统错误'
				})
				reject(error)
			}
		})
	})
})


export function getHost() {
	// 判断变量名是否在需要存储的数组中
	let port = uni.getSystemInfoSync().platform
	let h = "127.0.0.1"
	let p = ":8083/"
	switch (port) {
		case 'android':
			// console.log('Android'); //android
			h= "10.0.2.2"
			break;
		case 'ios':
			// console.log('iOS'); //ios
			h= "10.0.2.2"
			break;
		default:
			// console.log('小程序'); //devtools
			break;
	} 
	return h+p
}

export function getTime(timestamp) {
    var date = new Date(timestamp);//时间戳为10位需*1000，时间戳为13位的话不需乘1000
    let Y = date.getFullYear(),
        M = (date.getMonth() + 1 < 10 ? '0' + (date.getMonth() + 1) : date.getMonth() + 1),
        D = (date.getDate() < 10 ? '0' + (date.getDate()) : date.getDate()),
        h = (date.getHours() < 10 ? '0' + (date.getHours()) : date.getHours()),
        m = (date.getMinutes() < 10 ? '0' + (date.getMinutes()) : date.getMinutes()),
        s = (date.getSeconds() < 10 ? '0' + (date.getSeconds()) : date.getSeconds());
    // return Y + '-' + M + '-' + D + ' ' + h + ':' + m + ':' + s
    return M + '-' + D + ' ' + h + ':' + m 
}
export function concat_url(host,content){
	return "http://"+ host +"file/static/"+content;
}

export function removeById(arr, id){
   const requiredIndex = arr.findIndex(el => {
      return el.id === String(id);
   });
   if(requiredIndex === -1){
      return false;
   };
   return !!arr.splice(requiredIndex, 1);
};