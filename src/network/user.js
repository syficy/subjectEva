import {userRequest} from "./request";


export function loginRequest(data){
    return userRequest({
        url:'/student/login',
        method: 'post',
        headers: { 'Content-Type': 'application/json' },
        data,
    })
}


export function registerRequest(data){
    return userRequest({
        url:'/student/register',
        method: 'post',
        headers: { 'Content-Type': 'application/json' },
        data,
    })
}


export function changePdRequest(data){
    return userRequest({
        url:'/student/changePd',
        method: 'post',
        headers: { 'Content-Type': 'application/json' },
        data,
    })
}



export function changeInfoRequest(data){
    return userRequest({
        url:'/student/changeInfo',
        method: 'post',
        headers: { 'Content-Type': 'application/json' },
        data,
    })
}

/*
import instance from './server.js'

export const loginRequest = (data) => {
    return new Promise ( (reslove,reject) => {
        window.console.log("hello2")
        instance({
            method: 'post',
            url: "http://localhost:8088/student/login",
            headers: { 'Content-Type': 'application/json' },
            data,
        })
    })
}

export const registerRequest = (data) => {
    return new Promise ( (reslove,reject) => {
        window.console.log("hello2")
        instance({
            method: 'post',
            url: "http://localhost:8088/student/register",
            headers: { 'Content-Type': 'application/json' },
            data,
        }).then(res => {
            let body = res.data
            if (body.code == 200) { //操作成功
                reslove(body)
            }else{
                reject(body)
            }
        }).catch(err => {
            reject(err)
        })
    })
}

*/
