<template>
    <div>
        <NavBar>
            <div slot="center">推荐课程</div>
        </NavBar>


        <template>
            <div class="infinite-list-wrapper" style="overflow:auto">
                <div class="block">
                    <el-carousel  style="margin-top: 5%;height: 225px">
                        <el-carousel-item>
                            <div style="display:flex;align-items:center; justify-content:center;">
                                <img src="../../assets/images/p3.png" >
                            </div>
                        </el-carousel-item>
                        <el-carousel-item>
                            <div style="display:flex;align-items:center; justify-content:center;">
                                <img src="../../assets/images/p2.png" >
                            </div>
                        </el-carousel-item>
                        <el-carousel-item>
                            <div style="display:flex;align-items:center; justify-content:center;">
                                <img src="../../assets/images/p1.png" >
                            </div>
                        </el-carousel-item>
                    </el-carousel>
                </div>



                <ul
                        class="list"
                        v-infinite-scroll="load"
                        infinite-scroll-disabled="disabled"
                        style="list-style: none; margin: 0; padding: 0;"
                >

                    <li v-for="i in searchData.length" v-bind:key="i" class="list-item" >
                        <div style="position: center">
                            <el-card style="width:95%; margin-top:10%; margin-left:2.5%" class="box-card">
                                <span>{{searchData[i-1].Subject}}--{{searchData[i-1].Teacher}}</span>
                                <div class="block" style="margin-top:10%;">
                                    <span class="demonstration">评分</span>
                                    <el-rate disabled v-model="searchData[i-1].Rate" ></el-rate>
                                </div>
                                <el-collapse  accordion>
                                    <el-collapse-item title="查看评价" name="i">
                                        <div>{{searchData[i-1].Evtext}}</div>
                                    </el-collapse-item>
                                </el-collapse>
                                <!--<div>评价者： <span>{{searchData[i-1].StudentId}}</span></div>-->
                            </el-card>
                        </div>
                    </li>
                </ul>

                <p v-if="loading">加载中...</p>
                <p v-if="noMore">没有更多了</p>
            </div>
        </template>

    </div>
</template>

<script>
import NavBar from "../../components/common/navbar/NavBar";
import {HighRageRequest} from "../../network/eva";
/*import {getHomeMultidata} from "../../network/home";*/
export default {
    name: 'home',
    components:{
      NavBar
    },
    data(){
        return{
            count: 6,
            loading: false,
            value1:3,
            loading: false,
            searchData:[],
        }
    },
    mounted() {
        window.console.log(this.value)
        HighRageRequest().then(res=>{
            window.console.log("create")
            if (res.code==200){
            }
            this.searchData=res.data
        })
    },
    computed: {
        noMore () {
            return this.count >= 20
        },
        disabled () {
            return this.loading || this.noMore
        }
    },
    methods: {
        load () {
            this.loading = true
            setTimeout(() => {
                this.count += 2
                this.loading = false
            }, 2000)
        }
    },
    created() {
/*        //1.请求多个数据
        getHomeMultidata().then(res=>{
            window.console.log(res)
        })*/
    }
}
</script>

<!--scoped 防止影响到别的组建-->
<style scoped>
    .text {
        font-size: 14px;
    }

    .item {
        padding: 18px 0;
    }

    .box-card {
        width: 80%
    }

    .el-carousel__item h3 {
        color: #475669;
        font-size: 14px;
        opacity: 0.75;
        line-height: 150px;
        margin: 0;
    }

    .el-carousel__item:nth-child(2n) {
        background-color: #99a9bf;
    }

    .el-carousel__item:nth-child(2n+1) {
        background-color: #d3dce6;
    }

    .img{
        height:225px;
    }
</style>