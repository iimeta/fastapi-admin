import{u as ta,B as oa,p as ua,y as sa,i as na,z as ia,_ as da}from"./index.05fcd0e5.js";/* empty css               *//* empty css                *//* empty css               *//* empty css              *//* empty css              *//* empty css                *//* empty css               *//* empty css              *//* empty css               *//* empty css                *//* empty css               *//* empty css               *//* empty css                */import{c as Z,S as ra}from"./sortable.esm.734c0c44.js";/* empty css               *//* empty css                *//* empty css                *//* empty css               *//* empty css               */import{d as pa,r as ye,e as d,c as ee,w as be,bP as ca,B as n,C as b,aH as l,aG as t,aL as i,aM as r,bu as O,aD as k,aJ as E,aI as U,u as _,F as h,D as ma,aE as ae,bQ as $,g as _a,n as fa,aW as va,aK as ya,aF as ba,aS as ka,b2 as ga,bC as ha,bA as $a,bB as qa,b1 as wa,bD as Ca,bE as Va,b5 as Ia,bF as Da,ab as Fa,aU as xa,bi as Sa,bj as Ba,bl as za,bm as Ea,b4 as Ua,bG as Aa,ad as Na,aT as Ra,bH as La,bI as Pa,aV as Oa,bR as Ta,bS as Ka,bT as Ma,bN as ja,a_ as Ha,bJ as Qa}from"./arco.f6ea4e94.js";import{h as Ga,u as Ja}from"./vue.945ef37b.js";import{u as Wa}from"./loading.b2615842.js";import{q as T}from"./common.4fed7ae4.js";import{s as Xa,q as Ya,a as Za,b as el,c as al}from"./key.eb0cde5c.js";import{g as ll,c as tl}from"./app.e5492c05.js";import{q as ol}from"./model.abdfe31b.js";import ul from"./index.59d73180.js";import"./chart.6e5bd655.js";import"./base.87fcf6e2.js";/* empty css                *//* empty css                */const sl={class:"container"},nl={class:"action-icon"},il={class:"action-icon"},dl={id:"tableSetting"},rl={style:{"margin-right":"4px",cursor:"move"}},pl={class:"title"},cl={key:0},ml={key:1},_l={key:0},fl={key:1},vl={name:"AppKeyList"},yl=pa({...vl,setup(bl){const{proxy:F}=_a(),K=Ga(),ke=ye({type:"checkbox",showCheckedAll:!0,onlyCurrent:!1}),le=d([]);(async()=>{try{const{data:e}=await ll();le.value=e.items,K.query.app_id&&(s.value.app_id=Number(K.query.app_id))}catch{}})();const M=d([]);(async()=>{try{const{data:e}=await ol();M.value=e.items}catch{}})();const ge=async e=>{v(!0);try{await Xa(e),F.$message.success("\u5220\u9664\u6210\u529F"),B()}catch{}finally{v(!1)}},te=()=>({type:1,user_id:d(),app_id:d(),key:"",models:[],quota:d(),quota_expires_at:[],status:d(),created_at:[]}),{loading:he,setLoading:v}=Wa(!0),{t:p}=ta(),oe=d([]),s=d(te()),C=d([]),A=d([]),j=d("medium"),V=d([]),I=d(!0),ue=d(),x={current:1,pageSize:20,showTotal:!0,showPageSize:!0,pageSizeOptions:[20,50,100,500,1e3]},N=ye({...x}),$e=ee(()=>[{name:p("size.mini"),value:"mini"},{name:p("size.small"),value:"small"},{name:p("size.medium"),value:"medium"},{name:p("size.large"),value:"large"}]),se=ee(()=>[{title:p("key.columns.user_id"),dataIndex:"user_id",slotName:"user_id",align:"center",width:80},{title:p("key.columns.app_id"),dataIndex:"app_id",slotName:"app_id",align:"center",width:80},{title:p("key.columns.key"),dataIndex:"key",slotName:"key",align:"center",width:220},{title:p("key.columns.quota"),dataIndex:"quota",slotName:"quota",align:"center",ellipsis:!0,tooltip:!0},{title:p("key.columns.used_quota"),dataIndex:"used_quota",slotName:"used_quota",align:"center",ellipsis:!0,tooltip:!0},{title:p("key.columns.quota_expires_at"),dataIndex:"quota_expires_at",slotName:"quota_expires_at",align:"center",ellipsis:!0,tooltip:!0},{title:p("key.columns.app.models"),dataIndex:"model_names",slotName:"model_names",align:"center",ellipsis:!0,tooltip:!0},{title:p("key.columns.remark"),dataIndex:"remark",slotName:"remark",align:"center",ellipsis:!0,tooltip:!0},{title:p("key.columns.status"),dataIndex:"status",slotName:"status",align:"center",width:65},{title:p("key.columns.operations"),dataIndex:"operations",slotName:"operations",align:"center",width:170}]);localStorage.getItem("userRole")==="user"&&se.value.splice(0,1);const qe=ee(()=>[{label:p("key.dict.status.1"),value:1},{label:p("key.dict.status.2"),value:2}]),S=async(e={...x,type:1,app_id:K.query.app_id})=>{v(!0);try{const{data:o}=await Ya(e);oe.value=o.items,N.current=e.current,N.pageSize=e.pageSize,N.total=o.paging.total}catch{}finally{v(!1)}},B=()=>{S({...x,...s.value})},we=e=>{S({...x,...s.value,current:e})},Ce=e=>{x.pageSize=e,S({...x,...s.value})};S();const Ve=()=>{s.value=te()},Ie=async e=>{v(!0);try{await Za(e),F.$message.success("\u64CD\u4F5C\u6210\u529F"),B()}catch{}finally{v(!1)}},De=(e,o)=>{j.value=e},Fe=(e,o,m)=>{e?C.value.splice(m,0,o):C.value=A.value.filter(f=>f.dataIndex!==o.dataIndex)},ne=(e,o,m,f=!1)=>{const g=f?Z(e):e;return o>-1&&m>-1&&g.splice(o,1,g.splice(m,1,g[o]).pop()),g},xe=e=>{e&&fa(()=>{const o=document.getElementById("tableSetting");new ra(o,{onEnd(m){const{oldIndex:f,newIndex:g}=m;ne(C.value,f,g),ne(A.value,f,g)}})})};be(()=>se.value,e=>{C.value=Z(e),C.value.forEach((o,m)=>{o.checked=!0}),A.value=Z(C.value)},{deep:!0,immediate:!0});const z=d(!1),ie=d(),u=d({}),Se=e=>{u.value.quota=e*5e5},Be=async e=>{var o,m;v(!0);try{u.value.id=e.id,u.value.key=e.key,u.value.is_limit_quota=e.is_limit_quota,u.value.quota=e.quota,u.value.quota_expires_at=e.quota_expires_at,u.value.models=e.models,u.value.ip_whitelist=((o=e.ip_whitelist)==null?void 0:o.join(`
`))||"",u.value.ip_blacklist=((m=e.ip_blacklist)==null?void 0:m.join(`
`))||"",u.value.remark=e.remark,z.value=!0}catch{}finally{v(!1)}},ze=async e=>{var m;if(await((m=ie.value)==null?void 0:m.validate())){z.value=!0,e(!1);return}v(!0);try{const{data:f}=await tl(u.value);navigator.clipboard.writeText(f.key),va.success(p("app.success.key_config")),e(),S()}catch{}finally{v(!1)}},Ee=()=>{z.value=!1},Ue=e=>{V.value=e,I.value=!e.length},H=e=>{if(V.value.length===0)F.$message.info("\u8BF7\u9009\u62E9\u8981\u64CD\u4F5C\u7684\u6570\u636E");else{let o=`\u662F\u5426\u786E\u5B9A\u64CD\u4F5C\u6240\u9009\u7684${V.value.length}\u6761\u6570\u636E?`;switch(e.action){case"status":e.value===1?o=`\u662F\u5426\u786E\u5B9A\u542F\u7528\u6240\u9009\u7684${V.value.length}\u6761\u6570\u636E?`:o=`\u662F\u5426\u786E\u5B9A\u7981\u7528\u6240\u9009\u7684${V.value.length}\u6761\u6570\u636E?`;break;case"delete":o=`\u662F\u5426\u786E\u5B9A\u5220\u9664\u6240\u9009\u7684${V.value.length}\u6761\u6570\u636E?`;break}F.$modal.warning({title:"\u8B66\u544A",titleAlign:"start",content:o,hideCancel:!1,onOk:()=>{v(!0),e.ids=V.value,el(e).then(m=>{v(!1),F.$message.success("\u64CD\u4F5C\u6210\u529F"),B(),ue.value.selectAll(!1)})}})}},{copy:Ae,copied:de}=Ja(),Ne=async e=>{const{data:o}=await al({id:e});Ae(o.key)};be(de,()=>{de.value&&F.$message.success("\u590D\u5236\u6210\u529F")});const Q=d(!1),re=d(),Re=e=>{Q.value=!0,re.value=e},Le=()=>{Q.value=!1};return(e,o)=>{const m=oa,f=ya,g=ba,R=ka,c=ga,y=ha,G=$a,L=qa,pe=wa,Pe=Ca,J=Va,ce=Ia,me=Da,Oe=Fa,q=xa,_e=ua,fe=Sa,W=Ba,Te=sa,Ke=za,Me=Ea,je=na,He=ia,Qe=Ua,Ge=Aa,Je=Na,ve=Ra,We=La,Xe=Pa,Ye=Oa,D=Ta,Ze=Ka,ea=Ma,X=ja,aa=Ha,la=Qa,P=ca("permission");return n(),b("div",sl,[l(g,{class:"container-breadcrumb"},{default:t(()=>[l(f,null,{default:t(()=>[l(m)]),_:1}),l(f,null,{default:t(()=>[i(r(e.$t("menu.key.app")),1)]),_:1}),l(f,null,{default:t(()=>[i(r(e.$t("menu.key.app.list")),1)]),_:1})]),_:1}),l(la,{class:"general-card",bordered:!1,"header-style":{padding:"20px"},"body-style":{padding:"25px 20px 20px 20px"}},{default:t(()=>[l(J,null,{default:t(()=>[l(y,{flex:1},{default:t(()=>[l(ce,{model:s.value,"label-col-props":{span:5},"wrapper-col-props":{span:18},"label-align":"left"},{default:t(()=>[l(J,{gutter:16},{default:t(()=>[O((n(),k(y,{span:8},{default:t(()=>[l(c,{field:"user_id",label:e.$t("key.form.userId")},{default:t(()=>[l(R,{modelValue:s.value.user_id,"onUpdate:modelValue":o[0]||(o[0]=a=>s.value.user_id=a),placeholder:e.$t("key.form.userId.placeholder"),min:1,max:9999999999999,"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1})),[[P,["admin"]]]),O((n(),k(y,{span:8},{default:t(()=>[l(c,{field:"app_id",label:e.$t("key.form.appId")},{default:t(()=>[l(R,{modelValue:s.value.app_id,"onUpdate:modelValue":o[1]||(o[1]=a=>s.value.app_id=a),placeholder:e.$t("key.form.appId.placeholder"),min:1,max:9999999999999,"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1})),[[P,["admin"]]]),O((n(),k(y,{span:8},{default:t(()=>[l(c,{field:"app_id",label:e.$t("key.form.app")},{default:t(()=>[l(L,{modelValue:s.value.app_id,"onUpdate:modelValue":o[2]||(o[2]=a=>s.value.app_id=a),placeholder:e.$t("key.form.selectDefault"),"allow-search":"","allow-clear":""},{default:t(()=>[(n(!0),b(E,null,U(le.value,a=>(n(),k(G,{key:a.app_id,value:a.app_id,label:a.name},null,8,["value","label"]))),128))]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1})),[[P,["user"]]]),l(y,{span:8},{default:t(()=>[l(c,{field:"key",label:e.$t("key.form.appkey")},{default:t(()=>[l(pe,{modelValue:s.value.key,"onUpdate:modelValue":o[3]||(o[3]=a=>s.value.key=a),placeholder:e.$t("key.form.appkey.placeholder"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1}),O((n(),k(y,{span:8},{default:t(()=>[l(c,{field:"models",label:e.$t("key.form.app.models")},{default:t(()=>[l(L,{modelValue:s.value.models,"onUpdate:modelValue":o[4]||(o[4]=a=>s.value.models=a),placeholder:e.$t("key.form.selectDefault"),"max-tag-count":2,multiple:"","allow-search":"","allow-clear":""},{default:t(()=>[(n(!0),b(E,null,U(M.value,a=>(n(),k(G,{key:a.id,value:a.id,label:a.name},null,8,["value","label"]))),128))]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1})),[[P,["user"]]]),l(y,{span:8},{default:t(()=>[l(c,{field:"quota",label:e.$t("key.form.quota")},{default:t(()=>[l(R,{modelValue:s.value.quota,"onUpdate:modelValue":o[5]||(o[5]=a=>s.value.quota=a),placeholder:e.$t("key.form.quota.placeholder"),min:1e-6,max:9999999999999,"allow-clear":""},null,8,["modelValue","placeholder","min"])]),_:1},8,["label"])]),_:1}),l(y,{span:8},{default:t(()=>[l(c,{field:"status",label:e.$t("key.form.status")},{default:t(()=>[l(L,{modelValue:s.value.status,"onUpdate:modelValue":o[6]||(o[6]=a=>s.value.status=a),options:_(qe),placeholder:e.$t("key.form.selectDefault"),"allow-clear":""},null,8,["modelValue","options","placeholder"])]),_:1},8,["label"])]),_:1}),l(y,{span:8},{default:t(()=>[l(c,{field:"quota_expires_at",label:e.$t("key.form.quota_expires_at")},{default:t(()=>[l(Pe,{modelValue:s.value.quota_expires_at,"onUpdate:modelValue":o[7]||(o[7]=a=>s.value.quota_expires_at=a),style:{width:"100%"}},null,8,["modelValue"])]),_:1},8,["label"])]),_:1})]),_:1})]),_:1},8,["model"])]),_:1}),l(me,{style:{height:"84px"},direction:"vertical"}),l(y,{flex:"86px",style:{"text-align":"right"}},{default:t(()=>[l(fe,{direction:"vertical",size:18},{default:t(()=>[l(q,{type:"primary",onClick:B},{icon:t(()=>[l(Oe)]),default:t(()=>[i(" "+r(e.$t("key.form.search")),1)]),_:1}),l(q,{onClick:Ve},{icon:t(()=>[l(_e)]),default:t(()=>[i(" "+r(e.$t("key.form.reset")),1)]),_:1})]),_:1})]),_:1})]),_:1}),l(me,{style:{"margin-top":"0","margin-bottom":"16px"}}),l(J,{style:{"margin-bottom":"16px"}},{default:t(()=>[l(y,{span:12},{default:t(()=>[l(fe,null,{default:t(()=>[l(q,{type:"primary",status:"success",disabled:I.value,title:I.value?"\u8BF7\u9009\u62E9\u8981\u64CD\u4F5C\u7684\u6570\u636E":"",onClick:o[8]||(o[8]=a=>H({action:"status",value:1}))},{default:t(()=>[i(" \u542F\u7528 ")]),_:1},8,["disabled","title"]),l(q,{type:"primary",status:"danger",disabled:I.value,title:I.value?"\u8BF7\u9009\u62E9\u8981\u64CD\u4F5C\u7684\u6570\u636E":"",onClick:o[9]||(o[9]=a=>H({action:"status",value:2}))},{default:t(()=>[i(" \u7981\u7528 ")]),_:1},8,["disabled","title"]),l(q,{type:"primary",status:"danger",disabled:I.value,title:I.value?"\u8BF7\u9009\u62E9\u8981\u64CD\u4F5C\u7684\u6570\u636E":"",onClick:o[10]||(o[10]=a=>H({action:"delete"}))},{default:t(()=>[i(" \u5220\u9664 ")]),_:1},8,["disabled","title"])]),_:1})]),_:1}),l(y,{span:12,style:{display:"flex",height:"32px","align-items":"center","justify-content":"end"}},{default:t(()=>[l(W,{content:e.$t("actions.refresh")},{default:t(()=>[h("div",{class:"action-icon",onClick:B},[l(_e,{size:"18"})])]),_:1},8,["content"]),l(Me,{onSelect:De},{content:t(()=>[(n(!0),b(E,null,U(_($e),a=>(n(),k(Ke,{key:a.value,value:a.value,class:ma({active:a.value===j.value})},{default:t(()=>[h("span",null,r(a.name),1)]),_:2},1032,["value","class"]))),128))]),default:t(()=>[l(W,{content:e.$t("actions.density")},{default:t(()=>[h("div",nl,[l(Te,{size:"18"})])]),_:1},8,["content"])]),_:1}),l(W,{content:e.$t("actions.column_setting")},{default:t(()=>[l(Ge,{trigger:"click",position:"bl",onPopupVisibleChange:xe},{content:t(()=>[h("div",dl,[(n(!0),b(E,null,U(A.value,(a,w)=>(n(),b("div",{key:a.dataIndex,class:"setting"},[h("div",rl,[l(He)]),h("div",null,[l(Qe,{modelValue:a.checked,"onUpdate:modelValue":Y=>a.checked=Y,onChange:Y=>Fe(Y,a,w)},null,8,["modelValue","onUpdate:modelValue","onChange"])]),h("div",pl,r(a.title==="#"?"\u5E8F\u5217\u53F7":a.title),1)]))),128))])]),default:t(()=>[h("div",il,[l(je,{size:"18"})])]),_:1})]),_:1},8,["content"])]),_:1})]),_:1}),l(Xe,{ref_key:"tableRef",ref:ue,"row-key":"id",loading:_(he),pagination:N,columns:C.value,data:oe.value,bordered:!1,size:j.value,"row-selection":ke,onPageChange:we,onPageSizeChange:Ce,onSelectionChange:Ue},{key:t(({record:a})=>[i(r(a.key.substr(0,10)+a.key.substr(-10))+" ",1),l(Je,{class:"copy-btn",onClick:w=>Ne(a.id)},null,8,["onClick"])]),quota:t(({record:a})=>[a.is_limit_quota?(n(),b("span",cl,r(a.quota>0?`$${_(T)(a.quota)}`:a.quota<0?`-$${_(T)(-a.quota)}`:"$0.00"),1)):(n(),b("span",ml,r(e.$t("key.columns.quota.no_limit")),1))]),used_quota:t(({record:a})=>[i(" $"+r(a.used_quota>0?_(T)(a.used_quota):"0.00"),1)]),quota_expires_at:t(({record:a})=>[i(r(a.is_limit_quota&&a.quota_expires_at||"-"),1)]),model_names:t(({record:a})=>[a.model_names?(n(),b("span",_l,r(a.model_names.join(",")),1)):(n(),b("span",fl,r(e.$t("key.columns.app.models.no_limit")),1))]),remark:t(({record:a})=>[i(r(a.remark||"-"),1)]),status:t(({record:a})=>[l(ve,{modelValue:a.status,"onUpdate:modelValue":w=>a.status=w,"checked-value":1,"unchecked-value":2,onChange:w=>Ie({id:`${a.id}`,status:Number(`${a.status}`)})},null,8,["modelValue","onUpdate:modelValue","onChange"])]),operations:t(({record:a})=>[l(q,{type:"text",size:"small",onClick:w=>Re(a.id)},{default:t(()=>[i(r(e.$t("key.columns.operations.view")),1)]),_:2},1032,["onClick"]),l(q,{type:"text",size:"small",onClick:w=>Be(a)},{default:t(()=>[i(r(e.$t("key.columns.operations.update")),1)]),_:2},1032,["onClick"]),l(We,{content:"\u4F60\u786E\u5B9A\u8981\u5220\u9664\u5417?",onOk:w=>ge({id:`${a.id}`})},{default:t(()=>[l(q,{type:"text",size:"small"},{default:t(()=>[i(r(e.$t("key.columns.operations.delete")),1)]),_:1})]),_:2},1032,["onOk"])]),_:1},8,["loading","pagination","columns","data","size","row-selection"]),l(Ye,{title:e.$t("menu.key.detail"),"unmount-on-close":"","render-to-body":"",width:700,footer:!1,visible:Q.value,onCancel:Le},{default:t(()=>[l(ul,{id:re.value},null,8,["id"])]),_:1},8,["title","visible"]),l(aa,{visible:z.value,"onUpdate:visible":o[19]||(o[19]=a=>z.value=a),width:600,title:e.$t("app.form.title.keyConfig"),"ok-text":e.$t("app.button.save"),onCancel:Ee,onBeforeOk:ze},{default:t(()=>[l(ce,{ref_key:"formRef",ref:ie,model:u.value},{default:t(()=>[l(c,{field:"key",label:e.$t("app.label.key")},{default:t(()=>[l(pe,{modelValue:u.value.key,"onUpdate:modelValue":o[11]||(o[11]=a=>u.value.key=a),placeholder:e.$t("app.placeholder.key"),readonly:""},null,8,["modelValue","placeholder"])]),_:1},8,["label"]),l(c,{field:"is_limit_quota",label:e.$t("app.label.isLimitQuota")},{default:t(()=>[l(ve,{modelValue:u.value.is_limit_quota,"onUpdate:modelValue":o[12]||(o[12]=a=>u.value.is_limit_quota=a)},null,8,["modelValue"])]),_:1},8,["label"]),u.value.is_limit_quota?(n(),k(c,{key:0,field:"quota",label:e.$t("app.label.quota"),rules:[{required:!0,message:e.$t("app.error.quota.required")}]},{default:t(()=>[l(R,{modelValue:u.value.quota,"onUpdate:modelValue":o[13]||(o[13]=a=>u.value.quota=a),placeholder:e.$t("app.placeholder.quota"),precision:0,min:0,max:9999999999999,style:{"margin-right":"10px"}},null,8,["modelValue","placeholder"]),h("div",null," $"+r(u.value.quota?_(T)(u.value.quota):"0.00"),1)]),_:1},8,["label","rules"])):ae("",!0),u.value.is_limit_quota?(n(),k(c,{key:1},{default:t(()=>[l(Ze,{type:"button",onChange:Se},{default:t(()=>[l(D,{value:1},{default:t(()=>[i(" $1 ")]),_:1}),l(D,{value:5},{default:t(()=>[i(" $5 ")]),_:1}),l(D,{value:10},{default:t(()=>[i(" $10 ")]),_:1}),l(D,{value:20},{default:t(()=>[i(" $20 ")]),_:1}),l(D,{value:100},{default:t(()=>[i(" $100 ")]),_:1}),l(D,{value:500},{default:t(()=>[i(" $500 ")]),_:1}),l(D,{value:1e3},{default:t(()=>[i(" $1000 ")]),_:1})]),_:1},8,["onChange"])]),_:1})):ae("",!0),u.value.is_limit_quota?(n(),k(c,{key:2,field:"quota_expires_at",label:e.$t("app.label.quota_expires_at")},{default:t(()=>[l(ea,{modelValue:u.value.quota_expires_at,"onUpdate:modelValue":o[14]||(o[14]=a=>u.value.quota_expires_at=a),placeholder:e.$t("app.placeholder.quota_expires_at"),"time-picker-props":{defaultValue:"23:59:59"},"disabled-date":a=>_($)(a).isBefore(_($)()),style:{width:"100%"},"show-time":"",shortcuts:[{label:"1",value:()=>_($)().add(1,"day")},{label:"7",value:()=>_($)().add(7,"day")},{label:"15",value:()=>_($)().add(15,"day")},{label:"30",value:()=>_($)().add(30,"day")},{label:"90",value:()=>_($)().add(90,"day")},{label:"180",value:()=>_($)().add(180,"day")},{label:"365",value:()=>_($)().add(365,"day")}]},null,8,["modelValue","placeholder","disabled-date","shortcuts"])]),_:1},8,["label"])):ae("",!0),l(c,{field:"models",label:e.$t("app.label.models")},{default:t(()=>[l(L,{modelValue:u.value.models,"onUpdate:modelValue":o[15]||(o[15]=a=>u.value.models=a),placeholder:e.$t("app.placeholder.models"),"max-tag-count":3,multiple:"","allow-search":"","allow-clear":""},{default:t(()=>[(n(!0),b(E,null,U(M.value,a=>(n(),k(G,{key:a.id,value:a.id,label:a.name},null,8,["value","label"]))),128))]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label"]),l(c,{field:"ip_whitelist",label:e.$t("app.label.ip_whitelist")},{default:t(()=>[l(X,{modelValue:u.value.ip_whitelist,"onUpdate:modelValue":o[16]||(o[16]=a=>u.value.ip_whitelist=a),placeholder:e.$t("app.placeholder.ip_whitelist"),"auto-size":{minRows:5,maxRows:5}},null,8,["modelValue","placeholder"])]),_:1},8,["label"]),l(c,{field:"ip_blacklist",label:e.$t("app.label.ip_blacklist")},{default:t(()=>[l(X,{modelValue:u.value.ip_blacklist,"onUpdate:modelValue":o[17]||(o[17]=a=>u.value.ip_blacklist=a),placeholder:e.$t("app.placeholder.ip_blacklist"),"auto-size":{minRows:5,maxRows:5}},null,8,["modelValue","placeholder"])]),_:1},8,["label"]),l(c,{field:"remark",label:e.$t("app.placeholder.remark")},{default:t(()=>[l(X,{modelValue:u.value.remark,"onUpdate:modelValue":o[18]||(o[18]=a=>u.value.remark=a),placeholder:e.$t("app.placeholder.remark")},null,8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1},8,["model"])]),_:1},8,["visible","title","ok-text"])]),_:1})])}}});const et=da(yl,[["__scopeId","data-v-155f8559"]]);export{et as default};
