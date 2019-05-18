package templateStr

var FormVue  = `
<style lang="scss">@import "./css/style.scss";</style>
<template>
    <div id="{% .Module %}{% .Dir %}">
        <el-form :model="form" :rules="rules" label-width="160px">
		{% range $field := .Fields %}
			<el-form-item label="{% $field %}" prop="name">
                <el-input v-model="form.name" placeholder="请输入" style="width: 260px"></el-input>
            </el-form-item>
		{% end %}

            <el-form-item label="select组件" prop="name">
                <el-select v-model="form.name" placeholder="请选择">
                    <el-option
                        label="选择一"
                        value="1">
                    </el-option>
                    <el-option
                        label="选择二"
                        value="1">
                    </el-option>
                    <el-option
                        label="选择三"
                        value="3">
                    </el-option>
                </el-select>
            </el-form-item>

            <el-form-item label="checkbox组件" prop="userType">
                <el-checkbox-group
                    v-model="checkedCities">
                    <el-checkbox v-for="city in cities" :label="city" :key="city">{{city}}</el-checkbox>
                </el-checkbox-group>
            </el-form-item>

            <el-form-item label="时间组件" prop="time">
                <el-date-picker
                    v-model="form.time"
                    type="datetime"
                    placeholder="选择日期时间">
                </el-date-picker>
            </el-form-item>

            <el-form-item style="margin-top: 20px">
                <el-button type="primary" @click="">提交</el-button>
            </el-form-item>
        </el-form>
    </div>
</template>
<script>
    // import RichText from '@/components/RichText/RichText'

    export default {
        data() {
            return {
                form: {},
                rules: {},
                checkedCities: [],
                cities: ['心脏病护士', '打针护士', '心脏病护士', '打针护士']
            }
        },
        components: {

        }
    }
</script>

`
