<template>
  <div>
    <!--    搜索表单-->
    <div style="margin-bottom: 20px">
      <el-input style="width: 240px" placeholder="请输入图书名称" v-model="params.book_name"></el-input>
      <el-input style="width: 240px; margin-left: 5px" placeholder="请输入图书标准码" v-model="params.book_id"></el-input>
      <el-input style="width: 240px; margin-left: 5px" placeholder="请输入用户名称" v-model="params.user_name"></el-input>
      <el-button style="margin-left: 5px" type="primary" @click="load"><i class="el-icon-search"></i> 搜索</el-button>
      <el-button style="margin-left: 5px" type="warning" @click="reset"><i class="el-icon-refresh"></i> 重置</el-button>
    </div>

    <el-table :data="tableData" stripe row-key="id"  default-expand-all>
      <el-table-column prop="id" label="编号" width="80"></el-table-column>
      <el-table-column prop="Books.name" label="图书名称"></el-table-column>
      <el-table-column prop="book_id" label="图书标准码"></el-table-column>
      <el-table-column prop="Users.name" label="用户名称"></el-table-column>
      <el-table-column prop="user_id" label="会员码"></el-table-column>
      <el-table-column prop="Users.phone" label="用户联系方式"></el-table-column>
      <el-table-column prop="Books.score" label="所用积分"></el-table-column>
      <el-table-column prop="created_at" label="借出日期">
        <template v-slot="scope">
          {{ new Date(scope.row.created_at).toLocaleDateString().slice(0, 10) }}
        </template>
      </el-table-column>
      <el-table-column prop="days" label="借出天数"></el-table-column>
      <el-table-column prop="return_date" label="归还日期">
        <template v-slot="scope">
          {{ new Date(scope.row.return_date).toLocaleDateString().slice(0, 10) }}
        </template>
      </el-table-column>
      <el-table-column prop="note" label="状态提醒">
        <template v-slot="scope">
          <el-tag type="success" v-if="scope.row.note === '外租中'">{{ scope.row.note }}</el-tag>
          <el-tag type="warning" v-if="scope.row.note === '即将到期'">{{ scope.row.note }}</el-tag>
          <el-tag type="danger" v-if="scope.row.note === '已过期'">{{ scope.row.note }}</el-tag>
          <el-tag type="primary" v-if="scope.row.note === '已归还'">{{ scope.row.note }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="管理">
        <template v-slot="scope" >
          <el-button v-if="scope.row.note !== '已归还'" type="primary" @click="returnBooks(scope.row)" >还书</el-button>
        </template>
      </el-table-column>
      <el-table-column label="操作">
        <template v-slot="scope">
          <el-popconfirm
              style="margin-left: 5px"
              title="您确定删除这行数据吗？"
              @confirm="del(scope.row.id)">
            <el-button type="danger" slot="reference">删除</el-button>
          </el-popconfirm>
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
          :total="total">
      </el-pagination>
    </div>

  </div>
</template>

<script>
import request from "@/utils/request";
import Cookies from 'js-cookie'

export default {
  name: 'BorrowList',
  data() {
    return {
      admin: Cookies.get('admin') ? JSON.parse(Cookies.get('admin')) : {},
      tableData: [],
      total: 0,
      params: {
        page: 1,
        size: 10,
        user_name: '',
        book_name: '',
        book_id: '',
      }
    }
  },
  created() {
    this.load()
  },
  methods: {
    load() {
      request.get('/borrow_list', {
        params: this.params
      }).then(res => {
        if (res.code === '200') {
          this.tableData = res.data.list
          this.total = res.data.count
        }
      })
    },
    reset() {
      this.params = {
        page: 1,
        size: 10,
        user_name: '',
        book_name: '',
        book_id: ''
      }
      this.load()
    },
    handleCurrentChange(page) {
      // 点击分页按钮触发分页
      this.params.page = page
      this.load()
    },
    del(id) {
      request.delete("/borrow_delete" ,{
        params:{
          "id":id
        }
      }).then(res => {
        if (res.code === '200') {
          this.$notify.success('删除成功')
          this.load()
        } else {
          this.$notify.error(res.msg)
        }
      })
    },
    returnBooks(data) {
      console.log(data)
      request.post("/retur_create", data).then(res => {
        if (res.code === '200') {
          this.$notify.success('还书成功')
          this.load()
        } else {
          this.$notify.error(res.msg)
        }
      })
    }
  }
}
</script>

<style scoped>

</style>