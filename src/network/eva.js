import {evaRequest} from "./request";



export function insertRequest(data){
    return evaRequest({
        url:'/evaluate/insert',
        method: 'post',
        headers: { 'Content-Type': 'application/json' },
        data,
    })
}

export function searchTeacherRequest(data){
    return evaRequest({
        url:'/evaluate/queryTeacher',
        method: 'post',
        headers: { 'Content-Type': 'application/json' },
        data,
    })
}


export function HighRageRequest (){
    return evaRequest({
        url:'/evaluate/queryHighRageTeacher',
        headers: { 'Content-Type': 'application/json' },
        method: 'post',
    })
}

export function querySubjectRequest (){
    return evaRequest({
        url:'/evaluate/querySubject',
        headers: { 'Content-Type': 'application/json' },
        method: 'post',
    })
}

