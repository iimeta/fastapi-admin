import{u as ye,E as $e,o as we,x as Ce,I as ke,y as Ie,_ as Ve}from"./index.05df1f52.js";/* empty css               *//* empty css               *//* empty css              *//* empty css              *//* empty css                *//* empty css                *//* empty css               *//* empty css              *//* empty css                */import{c as T,S as ze}from"./sortable.esm.2109e0e3.js";/* empty css               *//* empty css               *//* empty css               *//* empty css               *//* empty css              */import{d as Se,r as j,e as b,c as C,w as De,B as v,C as $,aH as e,aG as a,aL as i,aM as d,u as k,F as p,aJ as G,aI as q,aD as Me,D as Ne,n as Be,aK as Fe,aF as Te,bC as Ae,b2 as Pe,bA as Ue,b1 as xe,bD as Le,bE as Ee,b5 as Oe,bF as je,ab as Ge,aU as qe,bi as He,a5 as Re,bj as Je,bl as Xe,bm as Ke,b4 as Qe,bG as We,bH as Ye,bI as Ze,bJ as et}from"./arco.aed15247.js";import{u as tt}from"./loading.b5911e1d.js";import{s as at,a as lt,b as ot}from"./model.89eea4c7.js";import"./chart.9aa6eafa.js";import"./vue.0cc5b64a.js";import"./base.87fcf6e2.js";const nt={class:"container"},st={class:"action-icon"},dt={class:"action-icon"},it={id:"tableSetting"},ct={style:{"margin-right":"4px",cursor:"move"}},ut={class:"title"},rt={key:0,class:"circle red"},mt={key:1,class:"circle"},pt={name:"ModelList"},_t=Se({...pt,setup(ft){const H=j({type:"checkbox",showCheckedAll:!0,onlyCurrent:!1}),R=async t=>{g(!0);try{await at(t),S()}catch{}finally{g(!1)}},A=()=>({corp:"",name:"",model:"",type:b(),status:b(),created_at:[]}),{loading:J,setLoading:g}=tt(!0),{t:n}=ye(),P=b([]),s=b(A()),_=b([]),I=b([]),D=b("medium"),h={current:1,pageSize:10,showTotal:!0,showPageSize:!0},V=j({...h}),X=C(()=>[{name:n("searchTable.size.mini"),value:"mini"},{name:n("searchTable.size.small"),value:"small"},{name:n("searchTable.size.medium"),value:"medium"},{name:n("searchTable.size.large"),value:"large"}]),K=C(()=>[{title:n("model.columns.corp"),dataIndex:"corp",slotName:"corp",align:"center",width:110},{title:n("model.columns.name"),dataIndex:"name",slotName:"name",align:"center"},{title:n("model.columns.model"),dataIndex:"model",slotName:"model",align:"center"},{title:n("model.columns.type"),dataIndex:"type",slotName:"type",align:"center"},{title:n("model.columns.billing_method"),dataIndex:"billing_method",slotName:"billingMethod",align:"center"},{title:n("model.columns.data_format"),dataIndex:"data_format",slotName:"dataFormat",align:"center"},{title:n("model.columns.status"),dataIndex:"status",slotName:"status",align:"center",width:80},{title:n("model.columns.updated_at"),dataIndex:"updated_at",slotName:"updated_at",align:"center",width:170},{title:n("model.columns.operations"),dataIndex:"operations",slotName:"operations",align:"center",width:220}]),Q=C(()=>[{label:n("model.dict.corp.OpenAI"),value:"OpenAI"},{label:n("model.dict.corp.Baidu"),value:"Baidu"},{label:n("model.dict.corp.Xfyun"),value:"Xfyun"},{label:n("model.dict.corp.Aliyun"),value:"Aliyun"},{label:n("model.dict.corp.GLM"),value:"GLM"},{label:n("model.dict.corp.Midjourney"),value:"Midjourney"}]),W=C(()=>[{label:n("model.dict.type.1"),value:1},{label:n("model.dict.type.2"),value:2},{label:n("model.dict.type.3"),value:3},{label:n("model.dict.type.4"),value:4}]),Y=C(()=>[{label:n("model.dict.status.1"),value:1},{label:n("model.dict.status.2"),value:2}]),z=async(t={...h})=>{g(!0);try{const{data:o}=await lt(t);P.value=o.items,V.current=t.current,V.pageSize=t.pageSize,V.total=o.paging.total}catch{}finally{g(!1)}},S=()=>{z({...h,...s.value})},Z=t=>{z({...h,...s.value,current:t})},ee=t=>{h.pageSize=t,z({...h,...s.value})};z();const te=()=>{s.value=A()},ae=async t=>{g(!0);try{t.status=t.status===1?2:1,await ot(t),S()}catch{}finally{g(!1)}},le=(t,o)=>{D.value=t},oe=(t,o,r)=>{t?_.value.splice(r,0,o):_.value=I.value.filter(c=>c.dataIndex!==o.dataIndex)},U=(t,o,r,c=!1)=>{const m=c?T(t):t;return o>-1&&r>-1&&m.splice(o,1,m.splice(r,1,m[o]).pop()),m},ne=t=>{t&&Be(()=>{const o=document.getElementById("tableSetting");new ze(o,{onEnd(r){const{oldIndex:c,newIndex:m}=r;U(_.value,c,m),U(I.value,c,m)}})})};return De(()=>K.value,t=>{_.value=T(t),_.value.forEach((o,r)=>{o.checked=!0}),I.value=T(_.value)},{deep:!0,immediate:!0}),(t,o)=>{const r=$e,c=Fe,m=Te,M=Ae,y=Pe,u=Ue,x=xe,se=Le,N=Ee,de=Oe,L=je,ie=Ge,f=qe,E=we,O=He,ce=Re,B=Je,ue=Ce,re=Xe,me=Ke,pe=ke,_e=Ie,fe=Qe,be=We,ve=Ye,ge=Ze,he=et;return v(),$("div",nt,[e(m,{class:"container-breadcrumb"},{default:a(()=>[e(c,null,{default:a(()=>[e(r)]),_:1}),e(c,null,{default:a(()=>[i(d(t.$t("menu.model")),1)]),_:1}),e(c,null,{default:a(()=>[i(d(t.$t("menu.model.list")),1)]),_:1})]),_:1}),e(he,{class:"general-card",title:t.$t("menu.model.list"),bordered:!1,"header-style":{padding:"20px"},"body-style":{padding:"0 20px 20px"}},{default:a(()=>[e(N,null,{default:a(()=>[e(u,{flex:1},{default:a(()=>[e(de,{model:s.value,"label-col-props":{span:5},"wrapper-col-props":{span:18},"label-align":"left"},{default:a(()=>[e(N,{gutter:16},{default:a(()=>[e(u,{span:8},{default:a(()=>[e(y,{field:"corp",label:t.$t("model.form.corp")},{default:a(()=>[e(M,{modelValue:s.value.corp,"onUpdate:modelValue":o[0]||(o[0]=l=>s.value.corp=l),options:k(Q),placeholder:t.$t("model.form.selectDefault"),"allow-clear":""},null,8,["modelValue","options","placeholder"])]),_:1},8,["label"])]),_:1}),e(u,{span:8},{default:a(()=>[e(y,{field:"name",label:t.$t("model.form.name")},{default:a(()=>[e(x,{modelValue:s.value.name,"onUpdate:modelValue":o[1]||(o[1]=l=>s.value.name=l),placeholder:t.$t("model.form.name.placeholder"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1}),e(u,{span:8},{default:a(()=>[e(y,{field:"model",label:t.$t("model.form.model")},{default:a(()=>[e(x,{modelValue:s.value.model,"onUpdate:modelValue":o[2]||(o[2]=l=>s.value.model=l),placeholder:t.$t("model.form.model.placeholder"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1}),e(u,{span:8},{default:a(()=>[e(y,{field:"type",label:t.$t("model.form.type")},{default:a(()=>[e(M,{modelValue:s.value.type,"onUpdate:modelValue":o[3]||(o[3]=l=>s.value.type=l),options:k(W),placeholder:t.$t("model.form.selectDefault"),"allow-clear":""},null,8,["modelValue","options","placeholder"])]),_:1},8,["label"])]),_:1}),e(u,{span:8},{default:a(()=>[e(y,{field:"status",label:t.$t("model.form.status")},{default:a(()=>[e(M,{modelValue:s.value.status,"onUpdate:modelValue":o[4]||(o[4]=l=>s.value.status=l),options:k(Y),placeholder:t.$t("model.form.selectDefault"),"allow-clear":""},null,8,["modelValue","options","placeholder"])]),_:1},8,["label"])]),_:1}),e(u,{span:8},{default:a(()=>[e(y,{field:"created_at",label:t.$t("model.form.created_at")},{default:a(()=>[e(se,{modelValue:s.value.created_at,"onUpdate:modelValue":o[5]||(o[5]=l=>s.value.created_at=l),style:{width:"100%"}},null,8,["modelValue"])]),_:1},8,["label"])]),_:1})]),_:1})]),_:1},8,["model"])]),_:1}),e(L,{style:{height:"84px"},direction:"vertical"}),e(u,{flex:"86px",style:{"text-align":"right"}},{default:a(()=>[e(O,{direction:"vertical",size:18},{default:a(()=>[e(f,{type:"primary",onClick:S},{icon:a(()=>[e(ie)]),default:a(()=>[i(" "+d(t.$t("model.form.search")),1)]),_:1}),e(f,{onClick:te},{icon:a(()=>[e(E)]),default:a(()=>[i(" "+d(t.$t("model.form.reset")),1)]),_:1})]),_:1})]),_:1})]),_:1}),e(L,{style:{"margin-top":"0"}}),e(N,{style:{"margin-bottom":"16px"}},{default:a(()=>[e(u,{span:12},{default:a(()=>[e(O,null,{default:a(()=>[e(f,{type:"primary",onClick:o[6]||(o[6]=l=>t.$router.push({name:"ModelCreate"}))},{icon:a(()=>[e(ce)]),default:a(()=>[i(" "+d(t.$t("model.operation.create")),1)]),_:1})]),_:1})]),_:1}),e(u,{span:12,style:{display:"flex","align-items":"center","justify-content":"end"}},{default:a(()=>[e(B,{content:t.$t("searchTable.actions.refresh")},{default:a(()=>[p("div",{class:"action-icon",onClick:S},[e(E,{size:"18"})])]),_:1},8,["content"]),e(me,{onSelect:le},{content:a(()=>[(v(!0),$(G,null,q(k(X),l=>(v(),Me(re,{key:l.value,value:l.value,class:Ne({active:l.value===D.value})},{default:a(()=>[p("span",null,d(l.name),1)]),_:2},1032,["value","class"]))),128))]),default:a(()=>[e(B,{content:t.$t("searchTable.actions.density")},{default:a(()=>[p("div",st,[e(ue,{size:"18"})])]),_:1},8,["content"])]),_:1}),e(B,{content:t.$t("searchTable.actions.columnSetting")},{default:a(()=>[e(be,{trigger:"click",position:"bl",onPopupVisibleChange:ne},{content:a(()=>[p("div",it,[(v(!0),$(G,null,q(I.value,(l,w)=>(v(),$("div",{key:l.dataIndex,class:"setting"},[p("div",ct,[e(_e)]),p("div",null,[e(fe,{modelValue:l.checked,"onUpdate:modelValue":F=>l.checked=F,onChange:F=>oe(F,l,w)},null,8,["modelValue","onUpdate:modelValue","onChange"])]),p("div",ut,d(l.title==="#"?"\u5E8F\u5217\u53F7":l.title),1)]))),128))])]),default:a(()=>[p("div",dt,[e(pe,{size:"18"})])]),_:1})]),_:1},8,["content"])]),_:1})]),_:1}),e(ge,{"row-key":"id",loading:k(J),pagination:V,columns:_.value,data:P.value,bordered:!1,size:D.value,"row-selection":H,onPageChange:Z,onPageSizeChange:ee},{type:a(({record:l})=>[i(d(t.$t(`model.dict.type.${l.type}`)),1)]),corp:a(({record:l})=>[i(d(t.$t(`model.dict.corp.${l.corp}`)),1)]),billingMethod:a(({record:l})=>[i(d(t.$t(`model.dict.billing_method.${l.billing_method}`)),1)]),dataFormat:a(({record:l})=>[i(d(t.$t(`model.dict.data_format.${l.data_format}`)),1)]),status:a(({record:l})=>[l.status===2?(v(),$("span",rt)):(v(),$("span",mt)),i(" "+d(t.$t(`model.dict.status.${l.status}`)),1)]),operations:a(({record:l})=>[e(f,{type:"text",size:"small",onClick:w=>t.$router.push({name:"ModelDetail",query:{id:`${l.id}`}})},{default:a(()=>[i(d(t.$t("model.columns.operations.view")),1)]),_:2},1032,["onClick"]),e(f,{type:"text",size:"small",onClick:w=>t.$router.push({name:"ModelUpdate",query:{id:`${l.id}`}})},{default:a(()=>[i(d(t.$t("model.columns.operations.update")),1)]),_:2},1032,["onClick"]),e(f,{type:"text",size:"small",onClick:w=>ae({id:`${l.id}`,status:Number(`${l.status}`)})},{default:a(()=>[i(d(t.$t(`model.columns.operations.status.${l.status}`)),1)]),_:2},1032,["onClick"]),e(ve,{content:"\u4F60\u786E\u5B9A\u8981\u5220\u9664\u5417?",onOk:w=>R({id:`${l.id}`})},{default:a(()=>[e(f,{type:"text",size:"small"},{default:a(()=>[i(d(t.$t("model.columns.operations.delete")),1)]),_:1})]),_:2},1032,["onOk"])]),_:1},8,["loading","pagination","columns","data","size","row-selection"])]),_:1},8,["title"])])}}});const xt=Ve(_t,[["__scopeId","data-v-d752aa76"]]);export{xt as default};
