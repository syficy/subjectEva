<template>
    <div class="tab-bar-item" @click="itemClick">
        <slot name="item-icon" v-if="!isActive"></slot>
        <slot name="item-icon-active" v-else></slot>
        <div :style="activeStyle">
            <slot name="item-text"></slot>
        </div>
    </div>
</template>

<script>
    export default {
        name: "TabBarItem.vue",
        props:{
            path: String,
            activeColor:{
                type:String,
                default: 'black'
            }
        },
        data(){
            return{
            }
        },
        computed:{
            isActive(){
                return this.$route.path.indexOf(this.path) !== -1
            },
            activeStyle(){
                return this.isActive ? {color:this.activeColor}:{}
            }
        },
        methods:{
            itemClick(){
                this.$router.replace(this.path)
            }
        }
    }

</script>

<style scoped>
    .tab-bar-item{

        flex:1;
        text-align: center;
        height: 49px;
    }

    .tab-bar-item i{
        height:1px;
        margin-bottem: 2px;
        margin-top: 5px;
    }

    .active{

    }
</style>