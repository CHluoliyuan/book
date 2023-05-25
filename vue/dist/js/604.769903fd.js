"use strict";(self["webpackChunkvue"]=self["webpackChunkvue"]||[]).push([[604],{4604:function(t,e,a){a.r(e),a.d(e,{default:function(){return u}});a(7658);var l=function(){var t=this,e=t._self._c;return e("div",[e("div",{staticStyle:{"margin-bottom":"20px"}},[e("el-input",{staticStyle:{width:"240px"},attrs:{placeholder:"请输入名称"},model:{value:t.params.name,callback:function(e){t.$set(t.params,"name",e)},expression:"params.name"}}),e("el-input",{staticStyle:{width:"240px","margin-left":"5px"},attrs:{placeholder:"请输入联系方式"},model:{value:t.params.phone,callback:function(e){t.$set(t.params,"phone",e)},expression:"params.phone"}}),e("el-button",{staticStyle:{"margin-left":"5px"},attrs:{type:"primary"},on:{click:t.load}},[e("i",{staticClass:"el-icon-search"}),t._v(" 搜索")]),e("el-button",{staticStyle:{"margin-left":"5px"},attrs:{type:"warning"},on:{click:t.reset}},[e("i",{staticClass:"el-icon-refresh"}),t._v(" 重置")])],1),e("el-table",{attrs:{data:t.tableData,stripe:""}},[e("el-table-column",{attrs:{prop:"id",label:"编号",width:"80"}}),e("el-table-column",{attrs:{prop:"name",label:"名称"}}),e("el-table-column",{attrs:{prop:"age",label:"年龄"}}),e("el-table-column",{attrs:{prop:"address",label:"地址"}}),e("el-table-column",{attrs:{prop:"phone",label:"联系方式"}}),e("el-table-column",{attrs:{prop:"sex",label:"性别"}}),e("el-table-column",{attrs:{prop:"account",label:"账户积分"}}),e("el-table-column",{attrs:{prop:"created_at",label:"创建时间"},scopedSlots:t._u([{key:"default",fn:function(e){return[t._v(" "+t._s(new Date(e.row.created_at).toLocaleDateString().slice(0,10))+" ")]}}])}),e("el-table-column",{attrs:{prop:"updated_at",label:"更新时间"},scopedSlots:t._u([{key:"default",fn:function(e){return[t._v(" "+t._s(new Date(e.row.updated_at).toLocaleDateString().slice(0,10))+" ")]}}])}),e("el-table-column",{attrs:{label:"操作",width:"230"},scopedSlots:t._u([{key:"default",fn:function(a){return[e("el-button",{attrs:{type:"warning"},on:{click:function(e){return t.handleAccountAdd(a.row)}}},[t._v("充值")]),e("el-button",{attrs:{type:"primary"},on:{click:function(e){return t.$router.push("/editUser?id="+a.row.id)}}},[t._v("编辑")]),e("el-popconfirm",{staticStyle:{"margin-left":"5px"},attrs:{title:"您确定删除这行数据吗？"},on:{confirm:function(e){return t.del(a.row.id)}}},[e("el-button",{attrs:{slot:"reference",type:"danger"},slot:"reference"},[t._v("删除")])],1)]}}])})],1),e("div",{staticStyle:{"margin-top":"20px"}},[e("el-pagination",{attrs:{background:"","current-page":t.params.page,"page-size":t.params.size,layout:"prev, pager, next",total:t.total},on:{"current-change":t.handleCurrentChange}})],1),e("el-dialog",{attrs:{title:"充值",visible:t.dialogFormVisible,width:"30%"},on:{"update:visible":function(e){t.dialogFormVisible=e}}},[e("el-form",{ref:"ruleForm",staticStyle:{width:"85%"},attrs:{model:t.form,"label-width":"100px",rules:t.rules}},[e("el-form-item",{attrs:{label:"当前账户积分",prop:"account"}},[e("el-input",{attrs:{disabled:"",autocomplete:"off"},model:{value:t.form.account,callback:function(e){t.$set(t.form,"account",e)},expression:"form.account"}})],1),e("el-form-item",{attrs:{label:"积分",prop:"score"}},[e("el-input",{attrs:{autocomplete:"off"},model:{value:t.form.score,callback:function(e){t.$set(t.form,"score",e)},expression:"form.score"}})],1)],1),e("div",{staticClass:"dialog-footer",attrs:{slot:"footer"},slot:"footer"},[e("el-button",{on:{click:function(e){t.dialogFormVisible=!1}}},[t._v("取 消")]),e("el-button",{attrs:{type:"primary"},on:{click:t.addAccount}},[t._v("确 定")])],1)],1)],1)},r=[],o=(a(1703),a(4471)),s={name:"User",data(){const t=(t,e,a)=>{e=parseInt(e),(e<10||e>200)&&a(new Error("请输入大于等于10小于或等于200的整数")),a()};return{tableData:[],total:0,params:{page:1,size:10,name:"",phone:""},dialogFormVisible:!1,form:{},rules:{score:[{required:!0,message:"请输入积分",trigger:"blur"},{validator:t,trigger:"blur"}]}}},created(){this.load()},methods:{load(){o.Z.get("/user_list",{params:this.params}).then((t=>{"200"===t.code&&(this.tableData=t.data.list,this.total=t.data.count)}))},reset(){this.params={page:1,size:10,name:"",phone:""},this.load()},handleCurrentChange(t){this.params.page=t,this.load()},del(t){o.Z["delete"]("/user_delete",{params:{id:t}}).then((t=>{"200"===t.code?(this.$notify.success("删除成功"),this.load()):this.$notify.error(t.msg)}))},handleAccountAdd(t){this.form=JSON.parse(JSON.stringify(t)),this.dialogFormVisible=!0},addAccount(){this.$refs["ruleForm"].validate((t=>{t&&o.Z.post("/user_account",this.form).then((t=>{"200"===t.code?(this.$notify.success("充值成功"),this.dialogFormVisible=!1,this.load()):this.$notify.error(t.msg)}))}))}}},n=s,i=a(3736),c=(0,i.Z)(n,l,r,!1,null,"6b141f68",null),u=c.exports}}]);
//# sourceMappingURL=604.769903fd.js.map