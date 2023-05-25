<template>
  <div>
    <!--    搜索表单-->
    <div style="margin-bottom: 20px">
      <el-input style="width: 240px" placeholder="请输入用户名" v-model="params.name"></el-input>
      <el-input style="width: 240px; margin-left: 5px" placeholder="请输入联系方式" v-model="params.phone"></el-input>
      <el-input style="width: 240px; margin-left: 5px" placeholder="请输入邮箱" v-model="params.email"></el-input>
      <el-button style="margin-left: 5px" type="primary" @click="load"><i class="el-icon-search"></i> 搜索</el-button>
      <el-button style="margin-left: 5px" type="warning" @click="reset"><i class="el-icon-refresh"></i> 重置</el-button>
    </div>

    <el-table :data="tableData" stripe>
      <el-table-column prop="id" label="编号" width="80"></el-table-column>
      <el-table-column prop="name" label="用户名"></el-table-column>
      <el-table-column prop="phone" label="联系方式"></el-table-column>
      <el-table-column prop="email" label="邮箱"></el-table-column>
      <el-table-column prop="created_at" label="创建时间">
        <template v-slot="scope">
          {{ new Date(scope.row.created_at).toLocaleDateString().slice(0, 10) }}
        </template>
      </el-table-column>
      <el-table-column prop="updated_at" label="更新时间">
        <template v-slot="scope">
          {{ new Date(scope.row.updated_at).toLocaleDateString().slice(0, 10) }}
        </template>
      </el-table-column>
    </el-table>

    <!--    分页-->
    <div style="margin-top: 20px">
      <el-pagination
          background
          :current-page="params.page"
          :page-size="params.size"
          layout="prev, pager, next"
          @current-change="handleCurrentChange"
          :total="count">
      </el-pagination>
    </div>

  </div>
</template>

<script>
import request from "@/utils/request";
export default {
  name: 'AdminList',
  data() {
    return {
      tableData: [],
      count: 0,
      params: {
        page: 1,
        size: 10,
        name: '',
        phone: '',
        email: ''
      },
    }
  },
  created() {
    this.load()
  },
  methods: {

    load() {
      request.get('/admin_list', {
        params: this.params
      }).then(res => {
        if (res.code === '200') {
          this.tableData = res.data.list
          this.count = res.data.count
        }
      })
    },
    reset() {
      this.params = {
        page: 1,
        size: 10,
        name: '',
        phone: '',
        email: ''
      }
      this.load()
    },
    handleCurrentChange(page) {
      // 点击分页按钮触发分页
      this.params.page = page
      this.load()
    },

  }
}
</script>

<style scoped>

</style>