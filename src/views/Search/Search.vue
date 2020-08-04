<template>
    <div>
        <nav-bar>
            <div slot="center">
                搜索
            </div>
        </nav-bar>

        <div class="block" style="margin-top: 10%">
            <el-select v-model="value" filterable placeholder="请选择">
                <el-option
                        v-for="item in options"
                        :key="item.value"
                        :label="item.label"
                        :value="item.value">
                </el-option>
            </el-select>
            <el-button icon="el-icon-search" circle style="margin-left: 10%" @click="search"></el-button>

<!--            <el-card style="width:100%; margin-top:40%" class="box-card">
                <span>课程+老师</span>
                <div class="block" style="margin-top:10%;">
                    <span class="demonstration">评分</span>
                    <el-rate v-model="value1" ></el-rate>
                </div>
                <el-collapse  accordion>
                    <el-collapse-item title="查看评价" name="i">
                        <div>与现实生活一致：与现实生活的流程、逻辑保持一致，遵循用户习惯的语言和概念；在界面中一致：所有的元素和结构需保持一致，比如：设计样式、图标和文本、元素的位置等。</div>
                    </el-collapse-item>
                </el-collapse>
            </el-card>-->
        </div>

        <div class="infinite-list-wrapper" style="overflow:auto">
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
                                <el-rate v-model="searchData[i-1].Rate" ></el-rate>
                            </div>
                            <el-collapse  accordion>
                                <el-collapse-item title="查看评价" name="i">
                                    <div>{{searchData[i-1].Evtext}}</div>
                                </el-collapse-item>
                            </el-collapse>
                        </el-card>
                    </div>
                </li>
            </ul>
            <p v-if="loading"></p>
            <p v-if="noMore">没有更多了</p>
        </div>













    </div>
</template>

<script>
    import NavBar from "../../components/common/navbar/NavBar";
    import {insertRequest, searchTeacherRequest,querySubjectRequest} from "../../network/eva";


    export default {
        name: "Search.vue",
        components:{
            NavBar
        },
        data(){
            return{
                options: [],
                value: '',
                count: 10,
                loading: false,
                searchData:[],
                subjectData:[],
            }
        },
        computed:{
            noMore () {
                return this.count >= 20
            },
            disabled () {
                return this.loading || this.noMore
            }
        },
        mounted() {
            querySubjectRequest().then(res=>{
                window.console.log(res.code)
                if (res.code==200){
                    window.console.log("查询subject成功")
                }
                this.subjectData=res.data
                var len = this.subjectData.length
                for(var i=0;i<len;i++)
                {
                    let data = {
                        "value":this.subjectData[i],
                        "label":this.subjectData[i]
                    }
                    this.options.push(data)
                }
                window.console.log(this.options)
            })
        },
        methods:{
            search(){
                window.console.log(this.value)
                let param ={
                    "subject": this.value
                }
                searchTeacherRequest(param).then(res=>{
                    window.console.log(res.code)
                    if (res.code==200){
                        this.$message.success("查询成功")
                    }
                    this.searchData=res.data
                    window.console.log(this.searchData)
                })

            },


            load () {
                this.loading = true
                setTimeout(() => {
                    this.count += 2
                    this.loading = false
                }, 2000)
            }
        }
    }
</script>

<style scoped>

</style>