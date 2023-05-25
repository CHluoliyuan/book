<template>
  <div>
    <!--    搜索表单-->
    <div style="margin-bottom: 20px">
      <el-input style="width: 240px" placeholder="请输入图书名称" v-model="params.name"></el-input>
      <el-input style="width: 240px; margin-left: 5px" placeholder="请输入图书标准码" v-model="params.id"></el-input>
      <el-button style="margin-left: 5px" type="primary" @click="load"><i class="el-icon-search"></i> 搜索</el-button>
      <el-button style="margin-left: 5px" type="warning" @click="reset"><i class="el-icon-refresh"></i> 重置</el-button>
    </div>

    <el-table :data="tableData" stripe row-key="id"  default-expand-all>
      <el-table-column prop="id" label="编号" width="80"></el-table-column>
      <el-table-column prop="name" label="图书名称"></el-table-column>
      <el-table-column prop="description" width="200" label="描述"></el-table-column>
      <el-table-column prop="publish_date" label="出版日期">
        <template v-slot="scope">
          {{ new Date(scope.row.publish_date).toLocaleDateString().slice(0, 10) }}
        </template>
      </el-table-column>
      <el-table-column prop="author" label="作者"></el-table-column>
      <el-table-column prop="publisher" label="出版社"></el-table-column>
      <el-table-column prop="category_id" label="分类">
        <template v-slot="scope">
          {{ categories.find(c => c.id === scope.row.category_id)?.name || '未知分类'}}

        </template>

      </el-table-column>

      <el-table-column prop="score" label="借书积分"></el-table-column>
      <el-table-column prop="nums" label="数量"></el-table-column>
      <el-table-column prop="cover" label="封面">
        <template v-slot="scope">
          <el-image :src="'http://47.108.85.66/api/'+scope.row.cover" :preview-src-list="['http://47.108.85.66/api/'+scope.row.cover]"></el-image>
        </template>
      </el-table-column>
      <el-table-column label="操作" width="140">
        <template v-slot="scope">
<!--          scope.row 就是当前行数据-->
          <el-button type="primary" @click="$router.push('/editBook?id=' + scope.row.id)">编辑</el-button>
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
  name: 'BookList',
  data() {
    return {
      categories:[],
      admin: Cookies.get('admin') ? JSON.parse(Cookies.get('admin')) : {},
      tableData: [],
      total: 0,
      params: {
        page: 1,
        size: 5,
        name: '',
        id: ''
      }
    }
  },
  created() {
    request.get('/category_list').then(res => {
      this.categories = res.data.list
    })

    this.load()
  },
  methods: {
    load() {
      request.get('/book_list', {
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
        name: '',
        id: ''
      }
      this.load()
    },
    handleCurrentChange(page) {
      // 点击分页按钮触发分页
      this.params.page = page
      this.load()
    },
    del(id) {
      request.delete("/book_delete",{
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
    }
  }
}
</script>

<style scoped>

</style>