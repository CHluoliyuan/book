<template>
  <div style="width: 80%">
    <div style="margin-bottom: 30px">编辑图书</div>
    <el-form :inline="true" :rules="rules" ref="ruleForm" :model="form" label-width="100px">
      <el-form-item label="名称" prop="name">
        <el-input v-model="form.name" placeholder="请输入名称"></el-input>
      </el-form-item>
      <el-form-item label="描述" prop="description">
        <el-input style="width: 400px" type="textarea" v-model="form.description" placeholder="请输入描述"></el-input>
      </el-form-item>
      <el-form-item label="出版日期" prop="publishDate">
        <el-date-picker
            v-model="form.publish_date"
            type="date"
            value-format="yyyy-MM-dd"
            placeholder="请选择出版日期">
        </el-date-picker>
      </el-form-item>
      <el-form-item label="作者" prop="author">
        <el-input v-model="form.author" placeholder="请输入作者"></el-input>
      </el-form-item>
      <el-form-item label="出版社" prop="publisher">
        <el-input v-model="form.publisher" placeholder="请输入出版社"></el-input>
      </el-form-item>
      <el-form-item label="分类" prop="category_id">
        <el-cascader
            style="width: 220px"
            :props="{ value: 'id', label: 'name'}"
            v-model="form.category_id"
            :options="categories"></el-cascader>
      </el-form-item>
      <el-form-item label="数量" prop="nums">
        <el-input v-model="form.nums" placeholder="数量"></el-input>
      </el-form-item>
      <el-form-item label="借书积分" prop="cover">
        <el-input-number v-model="form.score" :min="10" :max="30" label="所需积分"></el-input-number>
      </el-form-item>
      <br>
      <el-form-item label="封面" prop="cover">
        <el-upload
            class="avatar-uploader"
            :action="'http://47.108.85.66/api/bookimgup?token=' + this.admin.token"
            :show-file-list="false"
            :on-success="handleCoverSuccess">
          <img v-if="form.cover" :src="'http://47.108.85.66/api/'+form.cover" class="avatar" alt="">
          <i v-else class="el-icon-plus avatar-uploader-icon"></i>
        </el-upload>
      </el-form-item>
    </el-form>

    <div style="text-align: center; margin-top: 30px">
      <el-button type="primary" @click="save" size="medium">提交</el-button>
    </div>
  </div>
</template>

<script>
import request from "@/utils/request";
import Cookies from "js-cookie";

export default {
  name: 'EditBook',
  data() {
    return {
      admin: Cookies.get('admin') ? JSON.parse(Cookies.get('admin')) : {},
      form: { score: 10 },
      categories: [],
      rules: {
        name: [
          { required: true, message: '请输入图书名称', trigger: 'blur'}
        ],
        score: [
          { required: true, message: '请输入借书积分', trigger: 'blur'}
        ]
      }
    }
  },
  created() {
    request.get('/category_list').then(res => {
      this.categories = res.data.list
    })

    const id = this.$route.query.id
    request.get("/book_detail",{
      params:{
        "id":id
      }
    }).then(res => {
      this.form = res.data
      if (this.form.publish_date){
        this.form.publish_date=new Date(this.form.publish_date).toISOString().slice(0, 10)
      }
    })
  },
  methods: {
    handleCoverSuccess(res) {
      if (res.code === '200') {
        this.form.cover = res.data
      }
    },
    save() {
      console.log(this.form)
      request.put('/book_update', this.form).then(res => {
        if (res.code === '200') {
          this.$notify.success('更新成功')
          this.$router.push("/bookList")
        } else {
          this.$notify.error(res.msg)
        }
      })
    }
  }
}

</script>

<style>
.avatar-uploader .el-upload {
  border: 1px dashed #d9d9d9;
  border-radius: 6px;
  cursor: pointer;
  position: relative;
  overflow: hidden;
}
.avatar-uploader .el-upload:hover {
  border-color: #409EFF;
}
.avatar-uploader-icon {
  font-size: 28px;
  color: #8c939d;
  width: 178px;
  height: 178px;
  line-height: 178px;
  text-align: center;
}
.avatar {
  width: 178px;
  height: 178px;
  display: block;
}
</style>

