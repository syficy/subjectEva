import axios  from 'axios'


export function userRequest(config){
        //1.创建axios的实例
        const instance = axios.create({
            baseURL: 'http://47.95.227.20:8088',
            timeout: 5000
        })
        // 2.axios的拦截器
        // 2.1.请求拦截的作用
        instance.interceptors.request.use(config => {
            return config
        }, err => {
            // console.log(err);
        })

        // 2.2.响应拦截
        instance.interceptors.response.use(res => {
            return res.data
        }, err => {
            console.log(err);
        })
        return instance(config) //instance是一个promise
}




export function evaRequest(config){
    //1.创建axios的实例
    const instance = axios.create({
        baseURL: 'http://47.95.227.20:8089',
        timeout: 5000
    })
    // 2.axios的拦截器
    // 2.1.请求拦截的作用
    instance.interceptors.request.use(config => {
        return config
    }, err => {
        // console.log(err);
    })

    // 2.2.响应拦截
    instance.interceptors.response.use(res => {
        return res.data
    }, err => {
        console.log(err);
    })
    return instance(config) //instance是一个promise
}










/*equest({
    url:'/student/login'
}).then(res=>{

}).catch(err=>{
    windowconsole.log(err)
})*/

















/*
import axios  from 'axios'


export function request(config){
    return new Promise((resolve,reject)=>{
        //1.创建axios的实例
        const instance = axios.create({
            baseURL: 'http://localhost:8089',
            timeout: 5000
        })

        //发送网络请求
        instance(config).then(res=>{
            resolve(res)
        }).catch(err=>{
            reject(err)
            })

    })
}



request({
    url:'/student/login'
}).then(res=>{

}).catch(err=>{
    windowconsole.log(err)
    })*/
