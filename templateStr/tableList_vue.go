package templateStr

var TableList  = `
<style lang="scss">@import "./css/style.scss";</style>
<template>
    <div id="{% .Module %}{% .Dir %}">
        <ToolBar>
            <el-button type="primary" icon="el-icon-plus" size="small">添加</el-button>
            <div>
                <el-input
                        placeholder="搜索标题"
                        size="small"
                        style="width: 140px"
                        v-model="searchParams.postTitle"
                        clearable>
                </el-input>
                <el-select v-model="searchParams.postType" size="small" clearable placeholder="请选择分类"
                           style="width: 140px">
                    <el-option
                            v-for="(v,k) in $Cfg.postType"
                            :key="k"
                            :label="v"
                            :value="k">
                    </el-option>
                </el-select>
                <el-button type="success"  size="small" @click="">查询</el-button>
				<el-button type="warning" size="small" @click="">重置</el-button>
            </div>
        </ToolBar>
        <el-table
                :data="tableData"
                border
                style="width: 100%">
		{% range $field := .Fields %}
			<el-table-column
					prop="name"
					label="{% $field %}">
			</el-table-column>
		{% end %}
            <el-table-column
                    label="操作"
                    width="280">
                <template slot-scope="scope">
                    <el-button @click="handleClick(scope.row)" type="info" size="small"
                               >查看</el-button>
                    <el-button @click="handleClick(scope.row)" type="primary" size="small"
                               >编辑</el-button>
                    <el-button @click="handleClick(scope.row)" type="danger" size="small"
                               >删除</el-button>
                    <el-button @click="handleClick(scope.row)" type="success" size="small"
                    >其他</el-button>
                </template>
            </el-table-column>
        </el-table>
        <Pagination
            :params="searchParams"
            :requestFunc="requestFunc"
            @returnData="returnData"
        />
    </div>
</template>

<script>
    import ToolBar from '@/components/ToolBar.vue';
    import HelpHint from '@/components/HelpHint.vue';
    import Pagination from '@/components/Pagination.vue';
    import paginationMixin from './mixin/pagination.js'

    export default {
        mixins:[paginationMixin],
        data() {
            return {
                tableData: [{
                    name: 'Lorem ipsum dolor sit amet,',
                }, {
                    name: 'Lorem ipsum dolor sit amet,',
                }]
            }
        },
        methods: {
            handleClick(row) {
                this.$alert(row, '标题名称', {
                    confirmButtonText: '确定',
                    callback: action => {}
                });
            },
        },
        created() {

        },
        components: {
            ToolBar, HelpHint, Pagination
        }
    }
</script>

`
