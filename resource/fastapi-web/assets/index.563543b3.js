import{u as be,A as Ze,p as ea,y as aa,i as la,z as ta,_ as oa}from"./index.d6462cde.js";/* empty css                *//* empty css                *//* empty css                *//* empty css               *//* empty css               *//* empty css              *//* empty css                *//* empty css               *//* empty css               *//* empty css               *//* empty css              *//* empty css                *//* empty css               *//* empty css                *//* empty css               *//* empty css               *//* empty css                *//* empty css                *//* empty css               *//* empty css                *//* empty css                *//* empty css               *//* empty css               */import{d as ye,e as b,B as s,C as r,aH as e,aG as a,u as n,aD as c,aM as u,aL as f,bJ as sa,bK as na,bL as ua,bM as ia,bN as pa,r as fe,c as Y,w as da,bS as ra,bu as ve,aJ as Z,aI as ee,F as D,D as ma,aE as ae,bT as S,n as ca,aW as _a,aK as fa,aF as va,aS as ba,b2 as ya,bC as ha,b1 as ka,bA as ga,bB as $a,bU as wa,bD as qa,b5 as Va,bE as Ca,ab as xa,aU as Ia,bi as Da,bj as Sa,bl as za,bm as Ua,b4 as Aa,bF as Na,aT as La,bG as Ba,bH as Ra,aV as Ta,bP as Fa,bQ as Ma,bV as Pa,bW as Ea,bO as Ka,a_ as Oa,bI as ja,g as Ha}from"./arco.54c7388d.js";import{u as he}from"./loading.7321a6c2.js";import{q as L}from"./common.df364eef.js";import{q as Qa,s as Ga,a as Ja,b as Wa,c as Xa,d as Ya}from"./app.79615dea.js";import{c as le,S as Za}from"./sortable.esm.777e758f.js";import{q as el,a as al}from"./model.e72c4173.js";/* empty css                *//* empty css                */import"./chart.f14251fc.js";import"./vue.aa90ed69.js";import"./base.87fcf6e2.js";const ll={style:{margin:"10px 0 30px 10px"}},tl={key:1},ol={key:1},sl={key:1},nl={key:1},ul={key:1},il={key:1},pl={key:1,style:{"max-height":"220px",display:"block",overflow:"auto"}},dl={key:1,style:{"max-height":"110px",display:"block",overflow:"auto"}},rl={key:1,style:{"max-height":"110px",display:"block",overflow:"auto"}},ml={key:1,style:{"max-height":"110px",display:"block",overflow:"auto"}},cl={key:1},_l={key:1},fl={key:1},vl={name:"AppDetail"},bl=ye({...vl,props:{id:{type:String,default:""}},setup(te){const F=te,{t:y}=be(),{loading:p,setLoading:v}=he(!0),m=b({});return(async(B={id:F.id})=>{v(!0);try{const{data:A}=await Qa(B);m.value=A}catch{}finally{v(!1)}})(),(B,A)=>{const $=sa,g=na,h=ua,R=ia,_=pa;return s(),r("div",ll,[e(_,{column:2,bordered:"","value-style":{width:"350px",padding:"5px 8px 5px 20px"}},{default:a(()=>[e(h,{label:n(y)("common.app_id")},{default:a(()=>[n(p)?(s(),c(g,{key:0,animation:!0},{default:a(()=>[e($,{rows:1})]),_:1})):(s(),r("span",tl,u(m.value.app_id),1))]),_:1},8,["label"]),e(h,{label:n(y)("app.detail.label.name")},{default:a(()=>[n(p)?(s(),c(g,{key:0,animation:!0},{default:a(()=>[e($,{rows:1})]),_:1})):(s(),r("span",ol,u(m.value.name),1))]),_:1},8,["label"]),e(h,{label:n(y)("app.detail.label.is_limit_quota")},{default:a(()=>{var d;return[n(p)?(s(),c(g,{key:0,animation:!0},{default:a(()=>[e($,{rows:1})]),_:1})):(s(),r("span",sl,u(n(y)(`dict.${((d=m.value)==null?void 0:d.is_limit_quota)||!1}`)),1))]}),_:1},8,["label"]),e(h,{label:n(y)("app.detail.label.quota")},{default:a(()=>{var d;return[n(p)?(s(),c(g,{key:0,animation:!0},{default:a(()=>[e($,{rows:1})]),_:1})):(s(),r("span",nl,u((d=m.value)!=null&&d.is_limit_quota?m.value.quota>0?`$${n(L)(m.value.quota)}`:"$0.00":"\u4E0D\u9650"),1))]}),_:1},8,["label"]),e(h,{label:n(y)("app.detail.label.used_quota")},{default:a(()=>[n(p)?(s(),c(g,{key:0,animation:!0},{default:a(()=>[e($,{rows:1})]),_:1})):(s(),r("span",ul,u(m.value.used_quota>0?`$${n(L)(m.value.used_quota)}`:"$0.00"),1))]),_:1},8,["label"]),e(h,{label:n(y)("app.detail.label.quota_expires_at")},{default:a(()=>{var d;return[n(p)?(s(),c(g,{key:0,animation:!0},{default:a(()=>[e($,{rows:1})]),_:1})):(s(),r("span",il,u((d=m.value)!=null&&d.is_limit_quota&&m.value.quota_expires_at||"-"),1))]}),_:1},8,["label"]),e(h,{label:n(y)("app.detail.label.models"),span:2},{default:a(()=>{var d,w;return[n(p)?(s(),c(g,{key:0,animation:!0},{default:a(()=>[e($,{rows:1})]),_:1})):(s(),r("span",pl,u(((w=(d=m.value)==null?void 0:d.model_names)==null?void 0:w.join(`
`))||"-"),1))]}),_:1},8,["label"]),e(h,{label:n(y)("app.detail.label.ip_whitelist")},{default:a(()=>{var d,w;return[n(p)?(s(),c(g,{key:0,animation:!0},{default:a(()=>[e($,{rows:1})]),_:1})):(s(),r("span",dl,u(((w=(d=m.value)==null?void 0:d.ip_whitelist)==null?void 0:w.join(`
`))||"-"),1))]}),_:1},8,["label"]),e(h,{label:n(y)("app.detail.label.ip_blacklist")},{default:a(()=>{var d,w;return[n(p)?(s(),c(g,{key:0,animation:!0},{default:a(()=>[e($,{rows:1})]),_:1})):(s(),r("span",rl,u(((w=(d=m.value)==null?void 0:d.ip_blacklist)==null?void 0:w.join(`
`))||"-"),1))]}),_:1},8,["label"]),e(h,{label:n(y)("common.remark")},{default:a(()=>[n(p)?(s(),c(g,{key:0,animation:!0},{default:a(()=>[e($,{rows:1})]),_:1})):(s(),r("span",ml,u(m.value.remark||"-"),1))]),_:1},8,["label"]),e(h,{label:n(y)("common.status")},{default:a(()=>[n(p)?(s(),c(g,{key:0,animation:!0},{default:a(()=>[e($,{rows:1})]),_:1})):(s(),r("span",cl,[m.value.status===1?(s(),c(R,{key:0,color:"green"},{default:a(()=>[f(u(B.$t(`dict.status.${m.value.status}`)),1)]),_:1})):(s(),c(R,{key:1,color:"red"},{default:a(()=>[f(u(B.$t(`dict.status.${m.value.status}`)),1)]),_:1}))]))]),_:1},8,["label"]),e(h,{label:n(y)("common.created_at")},{default:a(()=>[n(p)?(s(),c(g,{key:0,animation:!0},{default:a(()=>[e($,{rows:1})]),_:1})):(s(),r("span",_l,u(m.value.created_at),1))]),_:1},8,["label"]),e(h,{label:n(y)("common.updated_at")},{default:a(()=>[n(p)?(s(),c(g,{key:0,animation:!0},{default:a(()=>[e($,{rows:1})]),_:1})):(s(),r("span",fl,u(m.value.updated_at),1))]),_:1},8,["label"])]),_:1})])}}}),yl={class:"container"},hl={class:"action-icon"},kl={class:"action-icon"},gl={id:"tableSetting"},$l={style:{"margin-right":"4px",cursor:"move"}},wl={class:"title"},ql={key:0},Vl={key:1},Cl={key:0},xl={key:1},Il={name:"AppList"},Dl=ye({...Il,setup(te){const{proxy:F}=Ha(),{loading:y,setLoading:p}=he(!0),{t:v}=be(),m=fe({type:"checkbox",showCheckedAll:!0,onlyCurrent:!1}),K=b([]);(async()=>{try{const{data:l}=await el();K.value=l.items}catch{}})();const A=b([]);(async()=>{p(!0);try{const{data:l}=await al();A.value=l.items}catch{}finally{p(!1)}})();const g=async l=>{p(!0);try{await Ga(l),F.$message.success("\u5220\u9664\u6210\u529F"),E()}catch{}finally{p(!1)}},h=()=>({user_id:b(),app_id:b(),name:"",models:[],app_key:"",type:b(),status:b(),quota_expires_at:[],created_at:[]}),R=b([]),_=b(h()),d=b([]),w=b([]),O=b("medium"),N={current:1,pageSize:20,showTotal:!0,showPageSize:!0,pageSizeOptions:[20,50,100,500,1e3]},M=fe({...N}),ke=Y(()=>[{name:v("size.mini"),value:"mini"},{name:v("size.small"),value:"small"},{name:v("size.medium"),value:"medium"},{name:v("size.large"),value:"large"}]),oe=Y(()=>[{title:v("app.columns.userId"),dataIndex:"user_id",slotName:"user_id",align:"center",width:80},{title:v("app.columns.appId"),dataIndex:"app_id",slotName:"app_id",align:"center",width:80},{title:v("app.columns.name"),dataIndex:"name",slotName:"name",align:"center",ellipsis:!0,tooltip:!0},{title:v("app.columns.models"),dataIndex:"model_names",slotName:"model_names",align:"center",ellipsis:!0,tooltip:!0},{title:v("app.columns.quota"),dataIndex:"quota",slotName:"quota",align:"center",ellipsis:!0,tooltip:!0},{title:v("app.columns.used_quota"),dataIndex:"used_quota",slotName:"used_quota",align:"center",ellipsis:!0,tooltip:!0},{title:v("app.columns.quota_expires_at"),dataIndex:"quota_expires_at",slotName:"quota_expires_at",align:"center",ellipsis:!0,tooltip:!0},{title:v("app.columns.status"),dataIndex:"status",slotName:"status",align:"center",width:65},{title:v("app.columns.operations"),dataIndex:"operations",slotName:"operations",align:"center",width:320}]);localStorage.getItem("userRole")==="user"&&oe.value.splice(0,1);const ge=Y(()=>[{label:v("app.dict.status.1"),value:1},{label:v("app.dict.status.2"),value:2}]),P=async(l={...N})=>{p(!0);try{const{data:o}=await Ja(l);R.value=o.items,M.current=l.current,M.pageSize=l.pageSize,M.total=o.paging.total}catch{}finally{p(!1)}},E=()=>{P({...N,..._.value})},$e=l=>{P({...N,..._.value,current:l})},we=l=>{N.pageSize=l,P({...N,..._.value})};P();const qe=()=>{_.value=h()},Ve=async l=>{p(!0);try{await Wa(l),F.$message.success("\u64CD\u4F5C\u6210\u529F"),E()}catch{}finally{p(!1)}},Ce=(l,o)=>{O.value=l},xe=(l,o,q)=>{l?d.value.splice(q,0,o):d.value=w.value.filter(V=>V.dataIndex!==o.dataIndex)},se=(l,o,q,V=!1)=>{const x=V?le(l):l;return o>-1&&q>-1&&x.splice(o,1,x.splice(q,1,x[o]).pop()),x},Ie=l=>{l&&ca(()=>{const o=document.getElementById("tableSetting");new Za(o,{onEnd(q){const{oldIndex:V,newIndex:x}=q;se(d.value,V,x),se(w.value,V,x)}})})};da(()=>oe.value,l=>{d.value=le(l),d.value.forEach((o,q)=>{o.checked=!0}),w.value=le(d.value)},{deep:!0,immediate:!0});const T=b(!1),ne=b(),i=b({}),De=l=>{i.value.quota=l*5e5},Se=async l=>{p(!0);try{const{data:o}=await Xa(l);i.value.app_id=o.app_id,i.value.key=o.key,T.value=!0}catch{}finally{p(!1)}},ze=async l=>{var q;if(await((q=ne.value)==null?void 0:q.validate())){T.value=!0,l(!1);return}p(!0);try{await Ya(i.value),navigator.clipboard.writeText(i.value.key),_a.success(v("app.success.key_config")),l()}catch{l(!1)}finally{p(!1)}},Ue=()=>{T.value=!1},j=b(!1),ue=b(),Ae=l=>{j.value=!0,ue.value=l},Ne=()=>{j.value=!1};return(l,o)=>{const q=Ze,V=fa,x=va,H=ba,k=ya,C=ha,Q=ka,Le=ga,ie=$a,Be=wa,G=qa,pe=Va,de=Ca,Re=xa,z=Ia,re=ea,me=Da,J=Sa,Te=aa,Fe=za,Me=Ua,Pe=la,Ee=ta,Ke=Aa,Oe=Na,ce=La,je=Ba,He=Ra,Qe=Ta,Ge=Fa,U=Ma,Je=Pa,We=Ea,W=Ka,Xe=Oa,Ye=ja,_e=ra("permission");return s(),r("div",yl,[e(x,{class:"container-breadcrumb"},{default:a(()=>[e(V,null,{default:a(()=>[e(q)]),_:1}),e(V,null,{default:a(()=>[f(u(l.$t("menu.app")),1)]),_:1}),e(V,null,{default:a(()=>[f(u(l.$t("menu.app.list")),1)]),_:1})]),_:1}),e(Ye,{class:"general-card",bordered:!1,"header-style":{padding:"20px"},"body-style":{padding:"25px 20px 20px 20px"}},{default:a(()=>[e(G,null,{default:a(()=>[e(C,{flex:1},{default:a(()=>[e(pe,{model:_.value,"label-col-props":{span:5},"wrapper-col-props":{span:18},"label-align":"left"},{default:a(()=>[e(G,{gutter:16},{default:a(()=>[ve((s(),c(C,{span:8},{default:a(()=>[e(k,{field:"user_id",label:l.$t("app.form.userId")},{default:a(()=>[e(H,{modelValue:_.value.user_id,"onUpdate:modelValue":o[0]||(o[0]=t=>_.value.user_id=t),placeholder:l.$t("app.form.userId.placeholder"),min:1,max:9999999999999,"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1})),[[_e,["admin"]]]),e(C,{span:8},{default:a(()=>[e(k,{field:"app_id",label:l.$t("app.form.appId")},{default:a(()=>[e(H,{modelValue:_.value.app_id,"onUpdate:modelValue":o[1]||(o[1]=t=>_.value.app_id=t),placeholder:l.$t("app.form.appId.placeholder"),min:1,max:9999999999999,"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1}),e(C,{span:8},{default:a(()=>[e(k,{field:"name",label:l.$t("app.form.name")},{default:a(()=>[e(Q,{modelValue:_.value.name,"onUpdate:modelValue":o[2]||(o[2]=t=>_.value.name=t),placeholder:l.$t("app.form.name.placeholder"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1}),ve((s(),c(C,{span:8},{default:a(()=>[e(k,{field:"models",label:l.$t("app.form.models")},{default:a(()=>[e(ie,{modelValue:_.value.models,"onUpdate:modelValue":o[3]||(o[3]=t=>_.value.models=t),placeholder:l.$t("app.form.selectDefault"),"max-tag-count":2,multiple:"","allow-search":"","allow-clear":""},{default:a(()=>[(s(!0),r(Z,null,ee(K.value,t=>(s(),c(Le,{key:t.id,value:t.id,label:t.name},null,8,["value","label"]))),128))]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1})),[[_e,["user"]]]),e(C,{span:8},{default:a(()=>[e(k,{field:"key",label:l.$t("app.form.key")},{default:a(()=>[e(Q,{modelValue:_.value.app_key,"onUpdate:modelValue":o[4]||(o[4]=t=>_.value.app_key=t),placeholder:l.$t("app.form.key.placeholder"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1}),e(C,{span:8},{default:a(()=>[e(k,{field:"status",label:l.$t("app.form.status")},{default:a(()=>[e(ie,{modelValue:_.value.status,"onUpdate:modelValue":o[5]||(o[5]=t=>_.value.status=t),options:n(ge),placeholder:l.$t("app.form.selectDefault"),"allow-clear":""},null,8,["modelValue","options","placeholder"])]),_:1},8,["label"])]),_:1}),e(C,{span:8},{default:a(()=>[e(k,{field:"quota_expires_at",label:l.$t("app.form.quota_expires_at")},{default:a(()=>[e(Be,{modelValue:_.value.quota_expires_at,"onUpdate:modelValue":o[6]||(o[6]=t=>_.value.quota_expires_at=t),style:{width:"100%"}},null,8,["modelValue"])]),_:1},8,["label"])]),_:1})]),_:1})]),_:1},8,["model"])]),_:1}),e(de,{style:{height:"84px"},direction:"vertical"}),e(C,{flex:"86px",style:{"text-align":"right"}},{default:a(()=>[e(me,{direction:"vertical",size:18},{default:a(()=>[e(z,{type:"primary",onClick:E},{icon:a(()=>[e(Re)]),default:a(()=>[f(" "+u(l.$t("app.form.search")),1)]),_:1}),e(z,{onClick:qe},{icon:a(()=>[e(re)]),default:a(()=>[f(" "+u(l.$t("app.form.reset")),1)]),_:1})]),_:1})]),_:1})]),_:1}),e(de,{style:{"margin-top":"0","margin-bottom":"16px"}}),e(G,{style:{"margin-bottom":"16px"}},{default:a(()=>[e(C,{span:12},{default:a(()=>[e(me,null,{default:a(()=>[e(z,{type:"primary",onClick:o[7]||(o[7]=t=>l.$router.push({name:"AppCreate"}))},{default:a(()=>[f(u(l.$t("app.operation.create")),1)]),_:1})]),_:1})]),_:1}),e(C,{span:12,style:{display:"flex",height:"32px","align-items":"center","justify-content":"end"}},{default:a(()=>[e(J,{content:l.$t("actions.refresh")},{default:a(()=>[D("div",{class:"action-icon",onClick:E},[e(re,{size:"18"})])]),_:1},8,["content"]),e(Me,{onSelect:Ce},{content:a(()=>[(s(!0),r(Z,null,ee(n(ke),t=>(s(),c(Fe,{key:t.value,value:t.value,class:ma({active:t.value===O.value})},{default:a(()=>[D("span",null,u(t.name),1)]),_:2},1032,["value","class"]))),128))]),default:a(()=>[e(J,{content:l.$t("actions.density")},{default:a(()=>[D("div",hl,[e(Te,{size:"18"})])]),_:1},8,["content"])]),_:1}),e(J,{content:l.$t("actions.column_setting")},{default:a(()=>[e(Oe,{trigger:"click",position:"bl",onPopupVisibleChange:Ie},{content:a(()=>[D("div",gl,[(s(!0),r(Z,null,ee(w.value,(t,I)=>(s(),r("div",{key:t.dataIndex,class:"setting"},[D("div",$l,[e(Ee)]),D("div",null,[e(Ke,{modelValue:t.checked,"onUpdate:modelValue":X=>t.checked=X,onChange:X=>xe(X,t,I)},null,8,["modelValue","onUpdate:modelValue","onChange"])]),D("div",wl,u(t.title==="#"?"\u5E8F\u5217\u53F7":t.title),1)]))),128))])]),default:a(()=>[D("div",kl,[e(Pe,{size:"18"})])]),_:1})]),_:1},8,["content"])]),_:1})]),_:1}),e(He,{"row-key":"id",loading:n(y),pagination:M,columns:d.value,data:R.value,bordered:!1,size:O.value,"row-selection":m,onPageChange:$e,onPageSizeChange:we},{model_names:a(({record:t})=>[t.model_names?(s(),r("span",ql,u(t.model_names.join(",")),1)):(s(),r("span",Vl,u(l.$t("app.columns.models.no_limit")),1))]),quota:a(({record:t})=>[t.is_limit_quota?(s(),r("span",Cl,u(t.quota>0?`$${n(L)(t.quota)}`:t.quota<0?`-$${n(L)(-t.quota)}`:"$0.00"),1)):(s(),r("span",xl,u(l.$t("app.columns.quota.no_limit")),1))]),used_quota:a(({record:t})=>[f(" $"+u(t.used_quota>0?n(L)(t.used_quota):"0.00"),1)]),quota_expires_at:a(({record:t})=>[f(u(t.is_limit_quota&&t.quota_expires_at||"-"),1)]),status:a(({record:t})=>[e(ce,{modelValue:t.status,"onUpdate:modelValue":I=>t.status=I,"checked-value":1,"unchecked-value":2,onChange:I=>Ve({id:`${t.id}`,status:Number(`${t.status}`)})},null,8,["modelValue","onUpdate:modelValue","onChange"])]),operations:a(({record:t})=>[e(z,{type:"text",size:"small",onClick:I=>Ae(t.id)},{default:a(()=>[f(u(l.$t("app.columns.operations.view")),1)]),_:2},1032,["onClick"]),e(z,{type:"text",size:"small",onClick:I=>l.$router.push({name:"AppUpdate",query:{id:`${t.id}`}})},{default:a(()=>[f(u(l.$t("app.columns.operations.update")),1)]),_:2},1032,["onClick"]),e(z,{type:"text",size:"small",onClick:I=>Se({app_id:`${t.app_id}`})},{default:a(()=>[f(u(l.$t("app.columns.operations.createKey")),1)]),_:2},1032,["onClick"]),e(z,{type:"text",size:"small",onClick:I=>l.$router.push({name:"AppKeyList",query:{app_id:`${t.app_id}`}})},{default:a(()=>[f(u(l.$t("app.columns.operations.manageKey")),1)]),_:2},1032,["onClick"]),e(je,{content:"\u4F60\u786E\u5B9A\u8981\u5220\u9664\u5417?",onOk:I=>g({id:`${t.id}`})},{default:a(()=>[e(z,{type:"text",size:"small"},{default:a(()=>[f(u(l.$t("app.columns.operations.delete")),1)]),_:1})]),_:2},1032,["onOk"])]),_:1},8,["loading","pagination","columns","data","size","row-selection"]),e(Qe,{title:l.$t("menu.app.detail"),"unmount-on-close":"","render-to-body":"",width:700,footer:!1,visible:j.value,onCancel:Ne},{default:a(()=>[e(bl,{id:ue.value},null,8,["id"])]),_:1},8,["title","visible"]),e(Xe,{visible:T.value,"onUpdate:visible":o[16]||(o[16]=t=>T.value=t),width:600,title:l.$t("app.form.title.keyConfig"),"ok-text":l.$t("app.button.save"),onCancel:Ue,onBeforeOk:ze},{default:a(()=>[e(pe,{ref_key:"formRef",ref:ne,model:i.value},{default:a(()=>[e(k,{field:"key",label:l.$t("app.label.key")},{default:a(()=>[e(Q,{modelValue:i.value.key,"onUpdate:modelValue":o[8]||(o[8]=t=>i.value.key=t),placeholder:l.$t("app.placeholder.key"),readonly:""},null,8,["modelValue","placeholder"])]),_:1},8,["label"]),e(k,{field:"models",label:l.$t("app.label.models")},{default:a(()=>[e(Ge,{modelValue:i.value.models,"onUpdate:modelValue":o[9]||(o[9]=t=>i.value.models=t),"allow-search":!0,"allow-clear":!0,"tree-checkable":!0,"tree-checked-strategy":"child",data:A.value,placeholder:l.$t("app.placeholder.key.models"),"max-tag-count":3},null,8,["modelValue","data","placeholder"])]),_:1},8,["label"]),e(k,{field:"is_limit_quota",label:l.$t("app.label.isLimitQuota")},{default:a(()=>[e(ce,{modelValue:i.value.is_limit_quota,"onUpdate:modelValue":o[10]||(o[10]=t=>i.value.is_limit_quota=t)},null,8,["modelValue"])]),_:1},8,["label"]),i.value.is_limit_quota?(s(),c(k,{key:0,field:"quota",label:l.$t("app.label.quota"),rules:[{required:!0,message:l.$t("app.error.quota.required")}]},{default:a(()=>[e(H,{modelValue:i.value.quota,"onUpdate:modelValue":o[11]||(o[11]=t=>i.value.quota=t),placeholder:l.$t("app.placeholder.quota"),precision:0,min:0,max:9999999999999,style:{"margin-right":"10px"}},null,8,["modelValue","placeholder"]),D("div",null," $"+u(i.value.quota?n(L)(i.value.quota):"0.00"),1)]),_:1},8,["label","rules"])):ae("",!0),i.value.is_limit_quota?(s(),c(k,{key:1},{default:a(()=>[e(Je,{type:"button",onChange:De},{default:a(()=>[e(U,{value:1},{default:a(()=>[f(" $1 ")]),_:1}),e(U,{value:5},{default:a(()=>[f(" $5 ")]),_:1}),e(U,{value:10},{default:a(()=>[f(" $10 ")]),_:1}),e(U,{value:20},{default:a(()=>[f(" $20 ")]),_:1}),e(U,{value:100},{default:a(()=>[f(" $100 ")]),_:1}),e(U,{value:500},{default:a(()=>[f(" $500 ")]),_:1}),e(U,{value:1e3},{default:a(()=>[f(" $1000 ")]),_:1})]),_:1},8,["onChange"])]),_:1})):ae("",!0),i.value.is_limit_quota?(s(),c(k,{key:2,field:"quota_expires_at",label:l.$t("app.label.quota_expires_at")},{default:a(()=>[e(We,{modelValue:i.value.quota_expires_at,"onUpdate:modelValue":o[12]||(o[12]=t=>i.value.quota_expires_at=t),placeholder:l.$t("app.placeholder.quota_expires_at"),"time-picker-props":{defaultValue:"23:59:59"},"disabled-date":t=>n(S)(t).isBefore(n(S)()),style:{width:"100%"},"show-time":"",shortcuts:[{label:"1",value:()=>n(S)().add(1,"day")},{label:"7",value:()=>n(S)().add(7,"day")},{label:"15",value:()=>n(S)().add(15,"day")},{label:"30",value:()=>n(S)().add(30,"day")},{label:"90",value:()=>n(S)().add(90,"day")},{label:"180",value:()=>n(S)().add(180,"day")},{label:"365",value:()=>n(S)().add(365,"day")}]},null,8,["modelValue","placeholder","disabled-date","shortcuts"])]),_:1},8,["label"])):ae("",!0),e(k,{field:"ip_whitelist",label:l.$t("app.label.ip_whitelist")},{default:a(()=>[e(W,{modelValue:i.value.ip_whitelist,"onUpdate:modelValue":o[13]||(o[13]=t=>i.value.ip_whitelist=t),placeholder:l.$t("app.placeholder.ip_whitelist"),"auto-size":{minRows:5,maxRows:5}},null,8,["modelValue","placeholder"])]),_:1},8,["label"]),e(k,{field:"ip_blacklist",label:l.$t("app.label.ip_blacklist")},{default:a(()=>[e(W,{modelValue:i.value.ip_blacklist,"onUpdate:modelValue":o[14]||(o[14]=t=>i.value.ip_blacklist=t),placeholder:l.$t("app.placeholder.ip_blacklist"),"auto-size":{minRows:5,maxRows:5}},null,8,["modelValue","placeholder"])]),_:1},8,["label"]),e(k,{field:"remark",label:l.$t("common.remark")},{default:a(()=>[e(W,{modelValue:i.value.remark,"onUpdate:modelValue":o[15]||(o[15]=t=>i.value.remark=t),placeholder:l.$t("app.placeholder.remark")},null,8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1},8,["model"])]),_:1},8,["visible","title","ok-text"])]),_:1})])}}});const mt=oa(Dl,[["__scopeId","data-v-3f77e90a"]]);export{mt as default};
