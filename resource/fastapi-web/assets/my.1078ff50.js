import{u as ve,E as be,o as ge,x as he,I as ye,y as we,_ as Ie}from"./index.e60636e1.js";/* empty css               *//* empty css               *//* empty css              *//* empty css              *//* empty css                *//* empty css                *//* empty css               *//* empty css              */import{c as D,S as $e}from"./sortable.esm.2109e0e3.js";/* empty css               *//* empty css               *//* empty css               *//* empty css               *//* empty css              */import{d as Ve,r as M,e as f,c as y,w as ke,B as v,C as h,aH as e,aG as a,aL as u,aM as d,u as w,F as p,aJ as O,aI as j,aD as Ce,D as Se,n as ze,aK as xe,aF as De,bC as Be,b2 as Te,bA as Ae,b1 as Ne,bD as Fe,bE as Ue,b5 as Pe,bF as Ee,ab as Le,aU as Me,bi as Oe,bj as je,bl as Re,bm as Ge,b4 as He,bG as Je,bI as Xe,bJ as Ze}from"./arco.aed15247.js";import{u as qe}from"./loading.b5911e1d.js";import{a as Ke}from"./model.3955e329.js";import"./chart.9aa6eafa.js";import"./vue.0cc5b64a.js";import"./base.87fcf6e2.js";const Qe={class:"container"},We={class:"action-icon"},Ye={class:"action-icon"},ea={id:"tableSetting"},aa={style:{"margin-right":"4px",cursor:"move"}},ta={class:"title"},la={key:0,class:"circle red"},oa={key:1,class:"circle"},na={name:"ModelList"},sa=Ve({...na,setup(da){const R=M({type:"checkbox",showCheckedAll:!0,onlyCurrent:!1}),B=()=>({corp:"",name:"",model:"",type:f(),status:f(),created_at:[]}),{loading:G,setLoading:T}=qe(!0),{t:n}=ve(),A=f([]),s=f(B()),_=f([]),I=f([]),k=f("medium"),b={current:1,pageSize:10,showTotal:!0,showPageSize:!0},$=M({...b}),H=y(()=>[{name:n("searchTable.size.mini"),value:"mini"},{name:n("searchTable.size.small"),value:"small"},{name:n("searchTable.size.medium"),value:"medium"},{name:n("searchTable.size.large"),value:"large"}]),J=y(()=>[{title:n("model.columns.corp"),dataIndex:"corp",slotName:"corp",align:"center",width:110},{title:n("model.columns.name"),dataIndex:"name",slotName:"name",align:"center"},{title:n("model.columns.model"),dataIndex:"model",slotName:"model",align:"center"},{title:n("model.columns.type"),dataIndex:"type",slotName:"type",align:"center"},{title:n("model.columns.prompt_price"),dataIndex:"prompt_price",slotName:"prompt_price",align:"center"},{title:n("model.columns.completion_price"),dataIndex:"completion_price",slotName:"completion_price",align:"center"},{title:n("model.columns.status"),dataIndex:"status",slotName:"status",align:"center",width:75},{title:n("model.columns.updated_at"),dataIndex:"updated_at",slotName:"updated_at",align:"center",width:132}]),X=y(()=>[{label:n("model.dict.corp.OpenAI"),value:"OpenAI"},{label:n("model.dict.corp.Baidu"),value:"Baidu"},{label:n("model.dict.corp.Xfyun"),value:"Xfyun"},{label:n("model.dict.corp.Aliyun"),value:"Aliyun"},{label:n("model.dict.corp.ZhipuAI"),value:"ZhipuAI"},{label:n("model.dict.corp.Midjourney"),value:"Midjourney"}]),Z=y(()=>[{label:n("model.dict.type.1"),value:1},{label:n("model.dict.type.2"),value:2},{label:n("model.dict.type.3"),value:3},{label:n("model.dict.type.4"),value:4}]),q=y(()=>[{label:n("model.dict.status.1"),value:1},{label:n("model.dict.status.2"),value:2}]),V=async(t={...b})=>{T(!0);try{const{data:o}=await Ke(t);A.value=o.items,$.current=t.current,$.pageSize=t.pageSize,$.total=o.paging.total}catch{}finally{T(!1)}},N=()=>{V({...b,...s.value})},K=t=>{V({...b,...s.value,current:t})},Q=t=>{b.pageSize=t,V({...b,...s.value})};V();const W=()=>{s.value=B()},Y=(t,o)=>{k.value=t},ee=(t,o,c)=>{t?_.value.splice(c,0,o):_.value=I.value.filter(i=>i.dataIndex!==o.dataIndex)},F=(t,o,c,i=!1)=>{const r=i?D(t):t;return o>-1&&c>-1&&r.splice(o,1,r.splice(c,1,r[o]).pop()),r},ae=t=>{t&&ze(()=>{const o=document.getElementById("tableSetting");new $e(o,{onEnd(c){const{oldIndex:i,newIndex:r}=c;F(_.value,i,r),F(I.value,i,r)}})})};return ke(()=>J.value,t=>{_.value=D(t),_.value.forEach((o,c)=>{o.checked=!0}),I.value=D(_.value)},{deep:!0,immediate:!0}),(t,o)=>{const c=be,i=xe,r=De,C=Be,g=Te,m=Ae,U=Ne,te=Fe,S=Ue,le=Pe,P=Ee,oe=Le,E=Me,L=ge,ne=Oe,z=je,se=he,de=Re,ce=Ge,ie=ye,re=we,me=He,ue=Je,pe=Xe,_e=Ze;return v(),h("div",Qe,[e(r,{class:"container-breadcrumb"},{default:a(()=>[e(i,null,{default:a(()=>[e(c)]),_:1}),e(i,null,{default:a(()=>[u(d(t.$t("menu.my.model")),1)]),_:1})]),_:1}),e(_e,{class:"general-card",bordered:!1,"header-style":{padding:"20px"},"body-style":{padding:"25px 20px 20px 20px"}},{default:a(()=>[e(S,null,{default:a(()=>[e(m,{flex:1},{default:a(()=>[e(le,{model:s.value,"label-col-props":{span:5},"wrapper-col-props":{span:18},"label-align":"left"},{default:a(()=>[e(S,{gutter:16},{default:a(()=>[e(m,{span:8},{default:a(()=>[e(g,{field:"corp",label:t.$t("model.form.corp")},{default:a(()=>[e(C,{modelValue:s.value.corp,"onUpdate:modelValue":o[0]||(o[0]=l=>s.value.corp=l),options:w(X),placeholder:t.$t("model.form.selectDefault"),"allow-clear":""},null,8,["modelValue","options","placeholder"])]),_:1},8,["label"])]),_:1}),e(m,{span:8},{default:a(()=>[e(g,{field:"name",label:t.$t("model.form.name")},{default:a(()=>[e(U,{modelValue:s.value.name,"onUpdate:modelValue":o[1]||(o[1]=l=>s.value.name=l),placeholder:t.$t("model.form.name.placeholder"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1}),e(m,{span:8},{default:a(()=>[e(g,{field:"model",label:t.$t("model.form.model")},{default:a(()=>[e(U,{modelValue:s.value.model,"onUpdate:modelValue":o[2]||(o[2]=l=>s.value.model=l),placeholder:t.$t("model.form.model.placeholder"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1}),e(m,{span:8},{default:a(()=>[e(g,{field:"type",label:t.$t("model.form.type")},{default:a(()=>[e(C,{modelValue:s.value.type,"onUpdate:modelValue":o[3]||(o[3]=l=>s.value.type=l),options:w(Z),placeholder:t.$t("model.form.selectDefault"),"allow-clear":""},null,8,["modelValue","options","placeholder"])]),_:1},8,["label"])]),_:1}),e(m,{span:8},{default:a(()=>[e(g,{field:"status",label:t.$t("model.form.status")},{default:a(()=>[e(C,{modelValue:s.value.status,"onUpdate:modelValue":o[4]||(o[4]=l=>s.value.status=l),options:w(q),placeholder:t.$t("model.form.selectDefault"),"allow-clear":""},null,8,["modelValue","options","placeholder"])]),_:1},8,["label"])]),_:1}),e(m,{span:8},{default:a(()=>[e(g,{field:"created_at",label:t.$t("model.form.created_at")},{default:a(()=>[e(te,{modelValue:s.value.created_at,"onUpdate:modelValue":o[5]||(o[5]=l=>s.value.created_at=l),style:{width:"100%"}},null,8,["modelValue"])]),_:1},8,["label"])]),_:1})]),_:1})]),_:1},8,["model"])]),_:1}),e(P,{style:{height:"84px"},direction:"vertical"}),e(m,{flex:"86px",style:{"text-align":"right"}},{default:a(()=>[e(ne,{direction:"vertical",size:18},{default:a(()=>[e(E,{type:"primary",onClick:N},{icon:a(()=>[e(oe)]),default:a(()=>[u(" "+d(t.$t("model.form.search")),1)]),_:1}),e(E,{onClick:W},{icon:a(()=>[e(L)]),default:a(()=>[u(" "+d(t.$t("model.form.reset")),1)]),_:1})]),_:1})]),_:1})]),_:1}),e(P,{style:{"margin-top":"0"}}),e(S,{style:{"margin-bottom":"16px"}},{default:a(()=>[e(m,{span:24,style:{display:"flex","align-items":"center","justify-content":"end"}},{default:a(()=>[e(z,{content:t.$t("searchTable.actions.refresh")},{default:a(()=>[p("div",{class:"action-icon",onClick:N},[e(L,{size:"18"})])]),_:1},8,["content"]),e(ce,{onSelect:Y},{content:a(()=>[(v(!0),h(O,null,j(w(H),l=>(v(),Ce(de,{key:l.value,value:l.value,class:Se({active:l.value===k.value})},{default:a(()=>[p("span",null,d(l.name),1)]),_:2},1032,["value","class"]))),128))]),default:a(()=>[e(z,{content:t.$t("searchTable.actions.density")},{default:a(()=>[p("div",We,[e(se,{size:"18"})])]),_:1},8,["content"])]),_:1}),e(z,{content:t.$t("searchTable.actions.columnSetting")},{default:a(()=>[e(ue,{trigger:"click",position:"bl",onPopupVisibleChange:ae},{content:a(()=>[p("div",ea,[(v(!0),h(O,null,j(I.value,(l,fe)=>(v(),h("div",{key:l.dataIndex,class:"setting"},[p("div",aa,[e(re)]),p("div",null,[e(me,{modelValue:l.checked,"onUpdate:modelValue":x=>l.checked=x,onChange:x=>ee(x,l,fe)},null,8,["modelValue","onUpdate:modelValue","onChange"])]),p("div",ta,d(l.title==="#"?"\u5E8F\u5217\u53F7":l.title),1)]))),128))])]),default:a(()=>[p("div",Ye,[e(ie,{size:"18"})])]),_:1})]),_:1},8,["content"])]),_:1})]),_:1}),e(pe,{"row-key":"id",loading:w(G),pagination:$,columns:_.value,data:A.value,bordered:!1,size:k.value,"row-selection":R,onPageChange:K,onPageSizeChange:Q},{type:a(({record:l})=>[u(d(t.$t(`model.dict.type.${l.type}`)),1)]),corp:a(({record:l})=>[u(d(t.$t(`model.dict.corp.${l.corp}`)),1)]),prompt_price:a(({record:l})=>[u(d(l.billing_method===1?`$${l.prompt_price}/k`:"-"),1)]),completion_price:a(({record:l})=>[u(d(l.billing_method===1?`$${l.completion_price}/k`:`$${l.fixed_price}/\u6B21`),1)]),status:a(({record:l})=>[l.status===2?(v(),h("span",la)):(v(),h("span",oa)),u(" "+d(t.$t(`model.dict.status.${l.status}`)),1)]),_:1},8,["loading","pagination","columns","data","size","row-selection"])]),_:1})])}}});const xa=Ie(sa,[["__scopeId","data-v-7a54dddd"]]);export{xa as default};