<template>
  <div style="width: 80%">
    <div style="margin-bottom: 30px">新增管理员</div>
    <el-form :inline="true" :model="form" :rules="rules" ref="ruleForm" label-width="100px">
      <el-form-item label="用户名" prop="name">
        <el-input v-model="form.name" placeholder="请输入姓名"></el-input>
      </el-form-item>
      <el-form-item label="联系方式" prop="phone">
        <el-input v-model="form.phone" placeholder="请输入联系方式"></el-input>
      </el-form-item>
      <el-form-item label="邮箱" prop="email">
        <el-input v-model="form.email" placeholder="请输入邮箱"></el-input>
      </el-form-item>
      <el-form-item label="验证码" prop="code">
        <el-input v-model="form.code" placeholder="请输入验证码">
        </el-input>
      </el-form-item>
      <el-button @click="code">获取验证码</el-button>
      <el-form-item label="密码" prop="password">
        <el-input v-model="form.password" placeholder="请输入密码"  type="password"></el-input>
      </el-form-item>
      <el-form-item label="确定密码" prop="password2">
        <el-input v-model="form.password2" placeholder="确认密码" type="password"></el-input>
      </el-form-item>
    </el-form>

    <div style="text-align: center; margin-top: 30px">
      <el-button type="primary" @click="save" size="medium">提交</el-button>
    </div>
  </div>
</template>

<script>
import request from "@/utils/request";

export default {
  name: 'AddAdmin',
  data() {
    const checkPhone = (rule, value, callback) => {
      if (!/^[1][3,4,5,6,7,8,9][0-9]{9}$/.test(value)) {
        callback(new Error('请输入合法的手机号'));
      }
      callback()
    };
    return {
      form: {},
      rules: {
        name: [
          { required: true, message: '请输入用户名', trigger: 'blur'},
          { min: 3, max: 10, message: '长度在3-10个字符', trigger: 'blur'}
        ],
        phone: [
          { validator: checkPhone, trigger: 'blur' }
        ]
      }
    }
  },
  methods: {
    code() {
      request.post('/code', this.form).then(res => {
            if (res.code === '200') {
              this.$notify.success(res.msg)
            } else {
              this.$notify.error(res.msg)
            }
          })
        },
    save() {
      this.$refs['ruleForm'].validate((valid) => {
        if (valid) {
          request.post('/register', this.form).then(res => {
            console.log(this.form)
            if (res.code === '200') {
              this.$notify.success('新增成功')
              this.$refs['ruleForm'].resetFields()
            } else {
              this.$notify.error(res.msg)
            }
          })
        }
      })
    },
    }

}

</script>

