<template>
  <div>
    <!--    搜索表单-->
    <div style="margin-bottom: 20px">
      <el-input style="width: 240px" placeholder="请输入分类名称" v-model="params.name"></el-input>
      <el-button style="margin-left: 5px" type="primary" @click="load"><i class="el-icon-search"></i> 搜索</el-button>
      <el-button style="margin-left: 5px" type="warning" @click="reset"><i class="el-icon-refresh"></i> 重置</el-button>
    </div>

    <el-table :data="tableData" stripe row-key="id"  default-expand-all>
      <el-table-column prop="id" label="编号" width="80"></el-table-column>
      <el-table-column prop="name" label="名称"></el-table-column>
      <el-table-column prop="remark" label="备注"></el-table-column>
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
      <el-table-column label="操作" width="280">
        <template v-slot="scope">
<!--          scope.row 就是当前行数据-->
          <el-button type="primary" @click="$router.push('/editCategory?id=' + scope.row.id)">编辑</el-button>
          <el-popconfirm
              style="margin-left: 5px"
              title="您确定删除这行数据吗？"
              @confirm="del(scope.row.id)"
          >
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
  name: 'CategoryList',
  data() {
    return {
      admin: Cookies.get('admin') ? JSON.parse(Cookies.get('admin')) : {},
      tableData: [],
      total: 0,
      dialogFormVisible: false,
      form: {},
      pid: null,
      params: {
        page: 1,
        size: 5,
        name: '',
      },
      rules: {
        name: [
          {required: true, message: '请输入分类名称', trigger: 'blur'}
        ]
      }
    }
  },
  created() {
    this.load()
  },
  methods: {
    load() {
      request.get('/category_list', {
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
        name: ''
      }
      this.load()
    },
    handleCurrentChange(page) {
      // 点击分页按钮触发分页
      this.params.page = page
      this.load()
    },
    del(id) {
      request.delete("/category_delete",{
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
  }
}
</script>

<style scoped>

</style>