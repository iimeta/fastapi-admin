import{u as ua,B as sa,p as na,y as ia,i as da,z as ra,_ as pa}from"./index.d6462cde.js";/* empty css                *//* empty css                *//* empty css                *//* empty css               *//* empty css               *//* empty css              *//* empty css                *//* empty css               *//* empty css               *//* empty css               *//* empty css              *//* empty css                *//* empty css               *//* empty css                *//* empty css               *//* empty css               *//* empty css                *//* empty css                *//* empty css               *//* empty css                *//* empty css                *//* empty css               *//* empty css               */import{d as ca,r as be,e as i,c as X,w as ke,bS as ma,B as d,C as b,aH as l,aG as t,aL as n,aM as r,bu as R,aD as k,aJ as L,aI as T,u as _,F as h,D as _a,aE as Y,bT as $,g as fa,n as va,aW as ya,aK as ba,aF as ka,aS as ga,b2 as ha,bC as $a,bA as qa,bB as wa,b1 as Ca,bU as Va,bD as Ia,b5 as Da,bE as Fa,ab as Sa,aU as xa,bi as Ba,bj as za,bl as Ea,bm as Ua,b4 as Aa,bF as Na,ad as Ra,aT as La,bG as Ta,bH as Ma,aV as Pa,bP as Oa,bQ as Ka,bV as ja,bW as Ha,bO as Qa,a_ as Ga,bI as Wa}from"./arco.54c7388d.js";import{h as Ja,u as Xa}from"./vue.aa90ed69.js";import{u as Ya}from"./loading.7321a6c2.js";import{q as M}from"./common.df364eef.js";import{s as Za,q as el,a as al,b as ll,c as tl}from"./key.432821f3.js";import{c as Z,S as ol}from"./sortable.esm.777e758f.js";import{g as ul,d as sl}from"./app.79615dea.js";import{q as nl,a as il}from"./model.e72c4173.js";import{_ as dl}from"./index.vue_vue_type_script_setup_true_lang.65e0ffe0.js";import"./chart.f14251fc.js";import"./base.87fcf6e2.js";/* empty css                *//* empty css                */const rl={class:"container"},pl={class:"action-icon"},cl={class:"action-icon"},ml={id:"tableSetting"},_l={style:{"margin-right":"4px",cursor:"move"}},fl={class:"title"},vl={key:0},yl={key:1},bl={key:0},kl={key:1},gl={name:"AppKeyList"},hl=ca({...gl,setup($l){const{proxy:F}=fa(),{loading:ge,setLoading:f}=Ya(!0),{t:p}=ua(),P=Ja(),he=be({type:"checkbox",showCheckedAll:!0,onlyCurrent:!1}),ee=i([]);(async()=>{try{const{data:e}=await ul();ee.value=e.items,P.query.app_id&&(s.value.app_id=Number(P.query.app_id))}catch{}})();const ae=i([]);(async()=>{try{const{data:e}=await nl();ae.value=e.items}catch{}})();const le=i([]);(async()=>{f(!0);try{const{data:e}=await il();le.value=e.items}catch{}finally{f(!1)}})();const $e=async e=>{f(!0);try{await Za(e),F.$message.success("\u5220\u9664\u6210\u529F"),B()}catch{}finally{f(!1)}},te=()=>({type:1,user_id:i(),app_id:i(),key:"",models:[],quota:i(),quota_expires_at:[],status:i(),created_at:[]}),oe=i([]),s=i(te()),C=i([]),E=i([]),O=i("medium"),V=i([]),I=i(!0),ue=i(),S={current:1,pageSize:20,showTotal:!0,showPageSize:!0,pageSizeOptions:[20,50,100,500,1e3]},U=be({...S}),qe=X(()=>[{name:p("size.mini"),value:"mini"},{name:p("size.small"),value:"small"},{name:p("size.medium"),value:"medium"},{name:p("size.large"),value:"large"}]),se=X(()=>[{title:p("key.columns.user_id"),dataIndex:"user_id",slotName:"user_id",align:"center",width:80},{title:p("key.columns.app_id"),dataIndex:"app_id",slotName:"app_id",align:"center",width:80},{title:p("key.columns.key"),dataIndex:"key",slotName:"key",align:"center",width:220},{title:p("key.columns.quota"),dataIndex:"quota",slotName:"quota",align:"center",ellipsis:!0,tooltip:!0},{title:p("key.columns.used_quota"),dataIndex:"used_quota",slotName:"used_quota",align:"center",ellipsis:!0,tooltip:!0},{title:p("key.columns.quota_expires_at"),dataIndex:"quota_expires_at",slotName:"quota_expires_at",align:"center",ellipsis:!0,tooltip:!0},{title:p("key.columns.app.models"),dataIndex:"model_names",slotName:"model_names",align:"center",ellipsis:!0,tooltip:!0},{title:p("key.columns.remark"),dataIndex:"remark",slotName:"remark",align:"center",ellipsis:!0,tooltip:!0},{title:p("key.columns.status"),dataIndex:"status",slotName:"status",align:"center",width:65},{title:p("key.columns.operations"),dataIndex:"operations",slotName:"operations",align:"center",width:170}]);localStorage.getItem("userRole")==="user"&&se.value.splice(0,1);const we=X(()=>[{label:p("key.dict.status.1"),value:1},{label:p("key.dict.status.2"),value:2}]),x=async(e={...S,type:1,app_id:P.query.app_id})=>{f(!0);try{const{data:o}=await el(e);oe.value=o.items,U.current=e.current,U.pageSize=e.pageSize,U.total=o.paging.total}catch{}finally{f(!1)}},B=()=>{x({...S,...s.value})},Ce=e=>{x({...S,...s.value,current:e})},Ve=e=>{S.pageSize=e,x({...S,...s.value})};x();const Ie=()=>{s.value=te()},De=async e=>{f(!0);try{await al(e),F.$message.success("\u64CD\u4F5C\u6210\u529F"),B()}catch{}finally{f(!1)}},Fe=(e,o)=>{O.value=e},Se=(e,o,m)=>{e?C.value.splice(m,0,o):C.value=E.value.filter(v=>v.dataIndex!==o.dataIndex)},ne=(e,o,m,v=!1)=>{const g=v?Z(e):e;return o>-1&&m>-1&&g.splice(o,1,g.splice(m,1,g[o]).pop()),g},xe=e=>{e&&va(()=>{const o=document.getElementById("tableSetting");new ol(o,{onEnd(m){const{oldIndex:v,newIndex:g}=m;ne(C.value,v,g),ne(E.value,v,g)}})})};ke(()=>se.value,e=>{C.value=Z(e),C.value.forEach((o,m)=>{o.checked=!0}),E.value=Z(C.value)},{deep:!0,immediate:!0});const z=i(!1),ie=i(),u=i({}),Be=e=>{u.value.quota=e*5e5},ze=async e=>{var o,m;f(!0);try{u.value.id=e.id,u.value.key=e.key,u.value.is_limit_quota=e.is_limit_quota,u.value.quota=e.quota,u.value.quota_expires_at=e.quota_expires_at,u.value.models=e.models,u.value.ip_whitelist=((o=e.ip_whitelist)==null?void 0:o.join(`
`))||"",u.value.ip_blacklist=((m=e.ip_blacklist)==null?void 0:m.join(`
`))||"",u.value.remark=e.remark,z.value=!0}catch{}finally{f(!1)}},Ee=async e=>{var m;if(await((m=ie.value)==null?void 0:m.validate())){z.value=!0,e(!1);return}f(!0);try{const{data:v}=await sl(u.value);navigator.clipboard.writeText(v.key),ya.success(p("app.success.key_config")),e(),x()}catch{e(!1)}finally{f(!1)}},Ue=()=>{z.value=!1},Ae=e=>{V.value=e,I.value=!e.length},K=e=>{if(V.value.length===0)F.$message.info("\u8BF7\u9009\u62E9\u8981\u64CD\u4F5C\u7684\u6570\u636E");else{let o=`\u662F\u5426\u786E\u5B9A\u64CD\u4F5C\u6240\u9009\u7684${V.value.length}\u6761\u6570\u636E?`;switch(e.action){case"status":e.value===1?o=`\u662F\u5426\u786E\u5B9A\u542F\u7528\u6240\u9009\u7684${V.value.length}\u6761\u6570\u636E?`:o=`\u662F\u5426\u786E\u5B9A\u7981\u7528\u6240\u9009\u7684${V.value.length}\u6761\u6570\u636E?`;break;case"delete":o=`\u662F\u5426\u786E\u5B9A\u5220\u9664\u6240\u9009\u7684${V.value.length}\u6761\u6570\u636E?`;break}F.$modal.warning({title:"\u8B66\u544A",titleAlign:"center",content:o,hideCancel:!1,onOk:()=>{f(!0),e.ids=V.value,ll(e).then(m=>{f(!1),F.$message.success("\u64CD\u4F5C\u6210\u529F"),B(),ue.value.selectAll(!1)})}})}},{copy:Ne,copied:de}=Xa(),Re=async e=>{const{data:o}=await tl({id:e});Ne(o.key)};ke(de,()=>{de.value&&F.$message.success("\u590D\u5236\u6210\u529F")});const j=i(!1),re=i(),Le=e=>{j.value=!0,re.value=e},Te=()=>{j.value=!1};return(e,o)=>{const m=sa,v=ba,g=ka,A=ga,c=ha,y=$a,pe=qa,H=wa,ce=Ca,Me=Va,Q=Ia,me=Da,_e=Fa,Pe=Sa,q=xa,fe=na,ve=Ba,G=za,Oe=ia,Ke=Ea,je=Ua,He=da,Qe=ra,Ge=Aa,We=Na,Je=Ra,ye=La,Xe=Ta,Ye=Ma,Ze=Pa,ea=Oa,D=Ka,aa=ja,la=Ha,W=Qa,ta=Ga,oa=Wa,N=ma("permission");return d(),b("div",rl,[l(g,{class:"container-breadcrumb"},{default:t(()=>[l(v,null,{default:t(()=>[l(m)]),_:1}),l(v,null,{default:t(()=>[n(r(e.$t("menu.key.app")),1)]),_:1}),l(v,null,{default:t(()=>[n(r(e.$t("menu.key.app.list")),1)]),_:1})]),_:1}),l(oa,{class:"general-card",bordered:!1,"header-style":{padding:"20px"},"body-style":{padding:"25px 20px 20px 20px"}},{default:t(()=>[l(Q,null,{default:t(()=>[l(y,{flex:1},{default:t(()=>[l(me,{model:s.value,"label-col-props":{span:5},"wrapper-col-props":{span:18},"label-align":"left"},{default:t(()=>[l(Q,{gutter:16},{default:t(()=>[R((d(),k(y,{span:8},{default:t(()=>[l(c,{field:"user_id",label:e.$t("key.form.userId")},{default:t(()=>[l(A,{modelValue:s.value.user_id,"onUpdate:modelValue":o[0]||(o[0]=a=>s.value.user_id=a),placeholder:e.$t("key.form.userId.placeholder"),min:1,max:9999999999999,"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1})),[[N,["admin"]]]),R((d(),k(y,{span:8},{default:t(()=>[l(c,{field:"app_id",label:e.$t("key.form.appId")},{default:t(()=>[l(A,{modelValue:s.value.app_id,"onUpdate:modelValue":o[1]||(o[1]=a=>s.value.app_id=a),placeholder:e.$t("key.form.appId.placeholder"),min:1,max:9999999999999,"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1})),[[N,["admin"]]]),R((d(),k(y,{span:8},{default:t(()=>[l(c,{field:"app_id",label:e.$t("key.form.app")},{default:t(()=>[l(H,{modelValue:s.value.app_id,"onUpdate:modelValue":o[2]||(o[2]=a=>s.value.app_id=a),placeholder:e.$t("key.form.selectDefault"),"allow-search":"","allow-clear":""},{default:t(()=>[(d(!0),b(L,null,T(ee.value,a=>(d(),k(pe,{key:a.app_id,value:a.app_id,label:a.name},null,8,["value","label"]))),128))]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1})),[[N,["user"]]]),l(y,{span:8},{default:t(()=>[l(c,{field:"key",label:e.$t("key.form.appkey")},{default:t(()=>[l(ce,{modelValue:s.value.key,"onUpdate:modelValue":o[3]||(o[3]=a=>s.value.key=a),placeholder:e.$t("key.form.appkey.placeholder"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1}),R((d(),k(y,{span:8},{default:t(()=>[l(c,{field:"models",label:e.$t("key.form.app.models")},{default:t(()=>[l(H,{modelValue:s.value.models,"onUpdate:modelValue":o[4]||(o[4]=a=>s.value.models=a),placeholder:e.$t("key.form.selectDefault"),"max-tag-count":2,multiple:"","allow-search":"","allow-clear":""},{default:t(()=>[(d(!0),b(L,null,T(ae.value,a=>(d(),k(pe,{key:a.id,value:a.id,label:a.name},null,8,["value","label"]))),128))]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1})),[[N,["user"]]]),l(y,{span:8},{default:t(()=>[l(c,{field:"quota",label:e.$t("key.form.quota")},{default:t(()=>[l(A,{modelValue:s.value.quota,"onUpdate:modelValue":o[5]||(o[5]=a=>s.value.quota=a),placeholder:e.$t("key.form.quota.placeholder"),min:1e-6,max:9999999999999,"allow-clear":""},null,8,["modelValue","placeholder","min"])]),_:1},8,["label"])]),_:1}),l(y,{span:8},{default:t(()=>[l(c,{field:"status",label:e.$t("key.form.status")},{default:t(()=>[l(H,{modelValue:s.value.status,"onUpdate:modelValue":o[6]||(o[6]=a=>s.value.status=a),options:_(we),placeholder:e.$t("key.form.selectDefault"),"allow-clear":""},null,8,["modelValue","options","placeholder"])]),_:1},8,["label"])]),_:1}),l(y,{span:8},{default:t(()=>[l(c,{field:"quota_expires_at",label:e.$t("key.form.quota_expires_at")},{default:t(()=>[l(Me,{modelValue:s.value.quota_expires_at,"onUpdate:modelValue":o[7]||(o[7]=a=>s.value.quota_expires_at=a),style:{width:"100%"}},null,8,["modelValue"])]),_:1},8,["label"])]),_:1})]),_:1})]),_:1},8,["model"])]),_:1}),l(_e,{style:{height:"84px"},direction:"vertical"}),l(y,{flex:"86px",style:{"text-align":"right"}},{default:t(()=>[l(ve,{direction:"vertical",size:18},{default:t(()=>[l(q,{type:"primary",onClick:B},{icon:t(()=>[l(Pe)]),default:t(()=>[n(" "+r(e.$t("key.form.search")),1)]),_:1}),l(q,{onClick:Ie},{icon:t(()=>[l(fe)]),default:t(()=>[n(" "+r(e.$t("key.form.reset")),1)]),_:1})]),_:1})]),_:1})]),_:1}),l(_e,{style:{"margin-top":"0","margin-bottom":"16px"}}),l(Q,{style:{"margin-bottom":"16px"}},{default:t(()=>[l(y,{span:12},{default:t(()=>[l(ve,null,{default:t(()=>[l(q,{type:"primary",status:"success",disabled:I.value,title:I.value?"\u8BF7\u9009\u62E9\u8981\u64CD\u4F5C\u7684\u6570\u636E":"",onClick:o[8]||(o[8]=a=>K({action:"status",value:1}))},{default:t(()=>[n(" \u542F\u7528 ")]),_:1},8,["disabled","title"]),l(q,{type:"primary",status:"danger",disabled:I.value,title:I.value?"\u8BF7\u9009\u62E9\u8981\u64CD\u4F5C\u7684\u6570\u636E":"",onClick:o[9]||(o[9]=a=>K({action:"status",value:2}))},{default:t(()=>[n(" \u7981\u7528 ")]),_:1},8,["disabled","title"]),l(q,{type:"primary",status:"danger",disabled:I.value,title:I.value?"\u8BF7\u9009\u62E9\u8981\u64CD\u4F5C\u7684\u6570\u636E":"",onClick:o[10]||(o[10]=a=>K({action:"delete"}))},{default:t(()=>[n(" \u5220\u9664 ")]),_:1},8,["disabled","title"])]),_:1})]),_:1}),l(y,{span:12,style:{display:"flex",height:"32px","align-items":"center","justify-content":"end"}},{default:t(()=>[l(G,{content:e.$t("actions.refresh")},{default:t(()=>[h("div",{class:"action-icon",onClick:B},[l(fe,{size:"18"})])]),_:1},8,["content"]),l(je,{onSelect:Fe},{content:t(()=>[(d(!0),b(L,null,T(_(qe),a=>(d(),k(Ke,{key:a.value,value:a.value,class:_a({active:a.value===O.value})},{default:t(()=>[h("span",null,r(a.name),1)]),_:2},1032,["value","class"]))),128))]),default:t(()=>[l(G,{content:e.$t("actions.density")},{default:t(()=>[h("div",pl,[l(Oe,{size:"18"})])]),_:1},8,["content"])]),_:1}),l(G,{content:e.$t("actions.column_setting")},{default:t(()=>[l(We,{trigger:"click",position:"bl",onPopupVisibleChange:xe},{content:t(()=>[h("div",ml,[(d(!0),b(L,null,T(E.value,(a,w)=>(d(),b("div",{key:a.dataIndex,class:"setting"},[h("div",_l,[l(Qe)]),h("div",null,[l(Ge,{modelValue:a.checked,"onUpdate:modelValue":J=>a.checked=J,onChange:J=>Se(J,a,w)},null,8,["modelValue","onUpdate:modelValue","onChange"])]),h("div",fl,r(a.title==="#"?"\u5E8F\u5217\u53F7":a.title),1)]))),128))])]),default:t(()=>[h("div",cl,[l(He,{size:"18"})])]),_:1})]),_:1},8,["content"])]),_:1})]),_:1}),l(Ye,{ref_key:"tableRef",ref:ue,"row-key":"id",loading:_(ge),pagination:U,columns:C.value,data:oe.value,bordered:!1,size:O.value,"row-selection":he,onPageChange:Ce,onPageSizeChange:Ve,onSelectionChange:Ae},{key:t(({record:a})=>[n(r(a.key.substr(0,10)+a.key.substr(-10))+" ",1),l(Je,{class:"copy-btn",onClick:w=>Re(a.id)},null,8,["onClick"])]),quota:t(({record:a})=>[a.is_limit_quota?(d(),b("span",vl,r(a.quota>0?`$${_(M)(a.quota)}`:a.quota<0?`-$${_(M)(-a.quota)}`:"$0.00"),1)):(d(),b("span",yl,r(e.$t("key.columns.quota.no_limit")),1))]),used_quota:t(({record:a})=>[n(" $"+r(a.used_quota>0?_(M)(a.used_quota):"0.00"),1)]),quota_expires_at:t(({record:a})=>[n(r(a.is_limit_quota&&a.quota_expires_at||"-"),1)]),model_names:t(({record:a})=>[a.model_names?(d(),b("span",bl,r(a.model_names.join(",")),1)):(d(),b("span",kl,r(e.$t("key.columns.app.models.no_limit")),1))]),remark:t(({record:a})=>[n(r(a.remark||"-"),1)]),status:t(({record:a})=>[l(ye,{modelValue:a.status,"onUpdate:modelValue":w=>a.status=w,"checked-value":1,"unchecked-value":2,onChange:w=>De({id:`${a.id}`,status:Number(`${a.status}`)})},null,8,["modelValue","onUpdate:modelValue","onChange"])]),operations:t(({record:a})=>[l(q,{type:"text",size:"small",onClick:w=>Le(a.id)},{default:t(()=>[n(r(e.$t("key.columns.operations.view")),1)]),_:2},1032,["onClick"]),l(q,{type:"text",size:"small",onClick:w=>ze(a)},{default:t(()=>[n(r(e.$t("key.columns.operations.update")),1)]),_:2},1032,["onClick"]),l(Xe,{content:"\u4F60\u786E\u5B9A\u8981\u5220\u9664\u5417?",onOk:w=>$e({id:`${a.id}`})},{default:t(()=>[l(q,{type:"text",size:"small"},{default:t(()=>[n(r(e.$t("key.columns.operations.delete")),1)]),_:1})]),_:2},1032,["onOk"])]),_:1},8,["loading","pagination","columns","data","size","row-selection"]),l(Ze,{title:e.$t("menu.key.detail"),"unmount-on-close":"","render-to-body":"",width:700,footer:!1,visible:j.value,onCancel:Te},{default:t(()=>[l(dl,{id:re.value},null,8,["id"])]),_:1},8,["title","visible"]),l(ta,{visible:z.value,"onUpdate:visible":o[19]||(o[19]=a=>z.value=a),width:600,title:e.$t("app.form.title.keyConfig"),"ok-text":e.$t("app.button.save"),onCancel:Ue,onBeforeOk:Ee},{default:t(()=>[l(me,{ref_key:"formRef",ref:ie,model:u.value},{default:t(()=>[l(c,{field:"key",label:e.$t("app.label.key")},{default:t(()=>[l(ce,{modelValue:u.value.key,"onUpdate:modelValue":o[11]||(o[11]=a=>u.value.key=a),placeholder:e.$t("app.placeholder.key"),readonly:""},null,8,["modelValue","placeholder"])]),_:1},8,["label"]),l(c,{field:"models",label:e.$t("app.label.models")},{default:t(()=>[l(ea,{modelValue:u.value.models,"onUpdate:modelValue":o[12]||(o[12]=a=>u.value.models=a),"allow-search":!0,"allow-clear":!0,"tree-checkable":!0,"tree-checked-strategy":"child",data:le.value,placeholder:e.$t("app.placeholder.key.models"),"max-tag-count":3},null,8,["modelValue","data","placeholder"])]),_:1},8,["label"]),l(c,{field:"is_limit_quota",label:e.$t("app.label.isLimitQuota")},{default:t(()=>[l(ye,{modelValue:u.value.is_limit_quota,"onUpdate:modelValue":o[13]||(o[13]=a=>u.value.is_limit_quota=a)},null,8,["modelValue"])]),_:1},8,["label"]),u.value.is_limit_quota?(d(),k(c,{key:0,field:"quota",label:e.$t("app.label.quota"),rules:[{required:!0,message:e.$t("app.error.quota.required")}]},{default:t(()=>[l(A,{modelValue:u.value.quota,"onUpdate:modelValue":o[14]||(o[14]=a=>u.value.quota=a),placeholder:e.$t("app.placeholder.quota"),precision:0,min:0,max:9999999999999,style:{"margin-right":"10px"}},null,8,["modelValue","placeholder"]),h("div",null," $"+r(u.value.quota?_(M)(u.value.quota):"0.00"),1)]),_:1},8,["label","rules"])):Y("",!0),u.value.is_limit_quota?(d(),k(c,{key:1},{default:t(()=>[l(aa,{type:"button",onChange:Be},{default:t(()=>[l(D,{value:1},{default:t(()=>[n(" $1 ")]),_:1}),l(D,{value:5},{default:t(()=>[n(" $5 ")]),_:1}),l(D,{value:10},{default:t(()=>[n(" $10 ")]),_:1}),l(D,{value:20},{default:t(()=>[n(" $20 ")]),_:1}),l(D,{value:100},{default:t(()=>[n(" $100 ")]),_:1}),l(D,{value:500},{default:t(()=>[n(" $500 ")]),_:1}),l(D,{value:1e3},{default:t(()=>[n(" $1000 ")]),_:1})]),_:1},8,["onChange"])]),_:1})):Y("",!0),u.value.is_limit_quota?(d(),k(c,{key:2,field:"quota_expires_at",label:e.$t("app.label.quota_expires_at")},{default:t(()=>[l(la,{modelValue:u.value.quota_expires_at,"onUpdate:modelValue":o[15]||(o[15]=a=>u.value.quota_expires_at=a),placeholder:e.$t("app.placeholder.quota_expires_at"),"time-picker-props":{defaultValue:"23:59:59"},"disabled-date":a=>_($)(a).isBefore(_($)()),style:{width:"100%"},"show-time":"",shortcuts:[{label:"1",value:()=>_($)().add(1,"day")},{label:"7",value:()=>_($)().add(7,"day")},{label:"15",value:()=>_($)().add(15,"day")},{label:"30",value:()=>_($)().add(30,"day")},{label:"90",value:()=>_($)().add(90,"day")},{label:"180",value:()=>_($)().add(180,"day")},{label:"365",value:()=>_($)().add(365,"day")}]},null,8,["modelValue","placeholder","disabled-date","shortcuts"])]),_:1},8,["label"])):Y("",!0),l(c,{field:"ip_whitelist",label:e.$t("app.label.ip_whitelist")},{default:t(()=>[l(W,{modelValue:u.value.ip_whitelist,"onUpdate:modelValue":o[16]||(o[16]=a=>u.value.ip_whitelist=a),placeholder:e.$t("app.placeholder.ip_whitelist"),"auto-size":{minRows:5,maxRows:5}},null,8,["modelValue","placeholder"])]),_:1},8,["label"]),l(c,{field:"ip_blacklist",label:e.$t("app.label.ip_blacklist")},{default:t(()=>[l(W,{modelValue:u.value.ip_blacklist,"onUpdate:modelValue":o[17]||(o[17]=a=>u.value.ip_blacklist=a),placeholder:e.$t("app.placeholder.ip_blacklist"),"auto-size":{minRows:5,maxRows:5}},null,8,["modelValue","placeholder"])]),_:1},8,["label"]),l(c,{field:"remark",label:e.$t("common.remark")},{default:t(()=>[l(W,{modelValue:u.value.remark,"onUpdate:modelValue":o[18]||(o[18]=a=>u.value.remark=a),placeholder:e.$t("app.placeholder.remark")},null,8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1},8,["model"])]),_:1},8,["visible","title","ok-text"])]),_:1})])}}});const rt=pa(hl,[["__scopeId","data-v-302c4c31"]]);export{rt as default};
