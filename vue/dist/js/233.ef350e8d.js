"use strict";(self["webpackChunkvue"]=self["webpackChunkvue"]||[]).push([[233],{5233:function(e,t,a){a.r(t),a.d(t,{default:function(){return u}});var l=function(){var e=this,t=e._self._c;return t("div",[t("div",{staticStyle:{"margin-bottom":"20px"}},[t("el-input",{staticStyle:{width:"240px"},attrs:{placeholder:"请输入图书名称"},model:{value:e.params.book_name,callback:function(t){e.$set(e.params,"book_name",t)},expression:"params.book_name"}}),t("el-input",{staticStyle:{width:"240px","margin-left":"5px"},attrs:{placeholder:"请输入图书标准码"},model:{value:e.params.book_id,callback:function(t){e.$set(e.params,"book_id",t)},expression:"params.book_id"}}),t("el-input",{staticStyle:{width:"240px","margin-left":"5px"},attrs:{placeholder:"请输入用户名称"},model:{value:e.params.user_name,callback:function(t){e.$set(e.params,"user_name",t)},expression:"params.user_name"}}),t("el-button",{staticStyle:{"margin-left":"5px"},attrs:{type:"primary"},on:{click:e.load}},[t("i",{staticClass:"el-icon-search"}),e._v(" 搜索")]),t("el-button",{staticStyle:{"margin-left":"5px"},attrs:{type:"warning"},on:{click:e.reset}},[t("i",{staticClass:"el-icon-refresh"}),e._v(" 重置")])],1),t("el-table",{attrs:{data:e.tableData,stripe:"","row-key":"id","default-expand-all":""}},[t("el-table-column",{attrs:{prop:"id",label:"编号",width:"80"}}),t("el-table-column",{attrs:{prop:"Books.name",label:"图书名称"}}),t("el-table-column",{attrs:{prop:"book_id",label:"图书标准码"}}),t("el-table-column",{attrs:{prop:"Users.name",label:"用户名称"}}),t("el-table-column",{attrs:{prop:"user_id",label:"会员码"}}),t("el-table-column",{attrs:{prop:"Users.phone",label:"用户联系方式"}}),t("el-table-column",{attrs:{prop:"Books.score",label:"所用积分"}}),t("el-table-column",{attrs:{prop:"borrow_date",label:"借出日期"},scopedSlots:e._u([{key:"default",fn:function(t){return[e._v(" "+e._s(new Date(t.row.borrow_date).toLocaleDateString().slice(0,10))+" ")]}}])}),t("el-table-column",{attrs:{prop:"days",label:"借出天数"}}),t("el-table-column",{attrs:{prop:"return_date",label:"归还日期"},scopedSlots:e._u([{key:"default",fn:function(t){return[e._v(" "+e._s(new Date(t.row.return_date).toLocaleDateString().slice(0,10))+" ")]}}])}),t("el-table-column",{attrs:{prop:"created_at",label:"实际归还日期"},scopedSlots:e._u([{key:"default",fn:function(t){return[e._v(" "+e._s(new Date(t.row.created_at).toLocaleDateString().slice(0,10))+" ")]}}])}),t("el-table-column",{attrs:{label:"操作"},scopedSlots:e._u([{key:"default",fn:function(a){return[t("el-popconfirm",{staticStyle:{"margin-left":"5px"},attrs:{title:"您确定删除这行数据吗？"},on:{confirm:function(t){return e.del(a.row.id)}}},[t("el-button",{attrs:{slot:"reference",type:"danger"},slot:"reference"},[e._v("删除")])],1)]}}])})],1),t("div",{staticStyle:{"margin-top":"20px"}},[t("el-pagination",{attrs:{background:"","current-page":e.params.page,"page-size":e.params.size,layout:"prev, pager, next",total:e.total},on:{"current-change":e.handleCurrentChange}})],1)],1)},r=[],o=a(4471),s=a(680),n={name:"ReturList",data(){return{admin:s.Z.get("admin")?JSON.parse(s.Z.get("admin")):{},tableData:[],total:0,params:{page:1,size:10,book_name:"",book_id:"",user_name:""}}},created(){this.load()},methods:{load(){o.Z.get("/retur_list",{params:this.params}).then((e=>{"200"===e.code&&(this.tableData=e.data.list,this.total=e.data.count),console.log(this.tableData)}))},reset(){this.params={page:1,size:10,book_name:"",book_id:"",user_name:""},this.load()},handleCurrentChange(e){this.params.page=e,this.load()},del(e){o.Z["delete"]("/retur_delete",{params:{id:e}}).then((e=>{"200"===e.code?(this.$notify.success("删除成功"),this.load()):this.$notify.error(e.msg)}))}}},i=n,p=a(3736),c=(0,p.Z)(i,l,r,!1,null,"1efb8c70",null),u=c.exports}}]);
//# sourceMappingURL=233.ef350e8d.js.map