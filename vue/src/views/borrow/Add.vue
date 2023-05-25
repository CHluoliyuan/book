<template>
  <div style="width: 80%">
    <div style="margin-bottom: 30px">新增借书记录</div>
    <el-form :inline="true" :rules="rules" ref="ruleForm" :model="form" label-width="100px">
      <el-form-item label="图书编号" prop="book_id">
        <el-select v-model="form.book_id" clearable filterable placeholder="请选择" @change="selBook">
          <el-option
              v-for="item in books"
              :key="item.id"
              :label="item.bookNo"
              :value="item.id">
          </el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="图书名称" prop="book_name">
        <el-input v-model="form.book_name" disabled></el-input>
      </el-form-item>
      <el-form-item label="所需积分" prop="book_score">
        <el-input v-model="form.book_score" disabled></el-input>
      </el-form-item>
      <el-form-item label="图书数量" prop="book_nums">
        <el-input v-model="form.book_nums" disabled></el-input>
      </el-form-item>
      <br />
      <el-form-item label="会员码" prop="user_id">
        <el-select v-model="form.user_id" filterable placeholder="请选择" @change="selUser">
          <el-option
              v-for="item in users"
              :key="item.id"
              :label="item.username"
              :value="item.id">
          </el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="用户名称" prop="user_name">
        <el-input disabled v-model="form.user_name"></el-input>
      </el-form-item>
      <el-form-item label="用户联系方式" prop="user_phone">
        <el-input disabled v-model="form.user_phone" ></el-input>
      </el-form-item>
      <el-form-item label="用户账户积分" prop="user_account">
        <el-input disabled v-model="form.user_account" ></el-input>
      </el-form-item>
      <el-form-item label="借出的天数" prop="days">
        <el-input-number v-model="form.days" :min="1" :max="30" label="借出的天数"></el-input-number>
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
  name: 'AddBook',
  data() {
    return {
      form: {days: 1},
      books: [],
      users: [],
      rules: {
        book_id: [
          { required: true, message: '请输入图书码', trigger: 'blur'}
        ],
        user_id: [
          { required: true, message: '请输入会员码', trigger: 'blur'}
        ]
      }
    }
  },
  created() {
    request.get('/book_list').then(res => {
      this.books = res.data.list
    })

    request.get('/user_list').then(res => {
      this.users = res.data.list
    })
  },
  methods: {
    save() {
      console.log(this.form)
      this.$refs['ruleForm'].validate((valid) => {
        if (valid) {
          request.post('/borrow_create', this.form).then(res => {
            if (res.code === '200') {
              this.$notify.success('新增成功')
              this.$refs['ruleForm'].resetFields()
              request.get('/user_list').then(res => {
                this.users = res.data.list
              })
              request.get('/book_list').then(res => {
                this.books = res.data.list
              })
            } else {
              this.$notify.error(res.msg)
            }
          })
        }
      })
    },
    selBook() {
      const book = this.books.find(v => v.id === this.form.book_id)
        this.form.book_name=book.name
        this.form.book_score = book.score
        this.form.book_nums = book.nums
    },
    selUser() {
      const user = this.users.find(v => v.id === this.form.user_id)
      this.form.user_name=user.name
      this.form.user_phone = user.phone
      this.form.user_account = user.account
    },
    }
  }

</script>

