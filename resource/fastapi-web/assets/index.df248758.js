import{u as Fe,E as Ie,p as De,y as Ee,i as Se,z as Be,_ as ze}from"./index.e0d6ab41.js";/* empty css                *//* empty css                *//* empty css               *//* empty css               *//* empty css              *//* empty css                *//* empty css               *//* empty css               *//* empty css                *//* empty css                *//* empty css               *//* empty css               *//* empty css                *//* empty css                *//* empty css               *//* empty css                *//* empty css                *//* empty css               *//* empty css               *//* empty css               *//* empty css              */import{d as xe,r as K,e as p,c as E,w as Ue,B as w,C as S,aH as e,aG as a,aL as s,aM as m,u as B,F as v,aJ as Q,aI as W,aD as Ae,D as Ne,n as Pe,aK as Le,aF as Oe,b1 as Te,b2 as Re,bC as He,bB as Me,bU as je,bD as qe,b5 as Ge,bE as Je,ab as Ke,aU as Qe,bi as We,bj as Xe,bl as Ye,bm as Ze,b4 as et,bF as tt,aT as at,bG as lt,bH as ot,bI as nt,g as ut}from"./arco.91d8d802.js";import{u as st}from"./loading.d8a03711.js";import{s as it,a as ct,b as rt,c as dt,d as pt}from"./corp.7ce589cd.js";import{c as L,S as mt}from"./sortable.esm.2649578f.js";import"./chart.1c4d013e.js";import"./vue.90059513.js";const _t={class:"container"},ft={class:"action-icon"},gt={class:"action-icon"},vt={id:"tableSetting"},bt={style:{"margin-right":"4px",cursor:"move"}},ht={class:"title"},Ct={name:"CorpList"},yt=xe({...Ct,setup(kt){const{proxy:y}=ut(),X=K({type:"checkbox",showCheckedAll:!0,onlyCurrent:!1}),Y=async t=>{i(!0);try{await it(t),y.$message.success("\u5220\u9664\u6210\u529F"),$()}catch{}finally{i(!1)}},O=()=>({name:"",code:"",remark:"",is_public:p(),status:p(),updated_at:[]}),{loading:Z,setLoading:i}=st(!0),{t:u}=Fe(),T=p([]),n=p(O()),b=p([]),F=p([]),z=p("medium"),h=p([]),C=p(!0),R=p(),k={current:1,pageSize:20,showTotal:!0,showPageSize:!0,pageSizeOptions:[20,50,100,500,1e3]},I=K({...k}),ee=E(()=>[{name:u("size.mini"),value:"mini"},{name:u("size.small"),value:"small"},{name:u("size.medium"),value:"medium"},{name:u("size.large"),value:"large"}]),te=E(()=>[{title:u("corp.columns.name"),dataIndex:"name",slotName:"name",align:"center"},{title:u("corp.columns.code"),dataIndex:"code",slotName:"code",align:"center"},{title:u("corp.columns.sort"),dataIndex:"sort",slotName:"sort",align:"center"},{title:u("corp.columns.is_public"),dataIndex:"is_public",slotName:"is_public",align:"center"},{title:u("corp.columns.remark"),dataIndex:"remark",slotName:"remark",align:"center"},{title:u("corp.columns.status"),dataIndex:"status",slotName:"status",align:"center"},{title:u("corp.columns.updated_at"),dataIndex:"updated_at",slotName:"updated_at",align:"center"},{title:u("corp.columns.operations"),dataIndex:"operations",slotName:"operations",align:"center",width:130}]),ae=E(()=>[{label:u("corp.dict.status.1"),value:1},{label:u("corp.dict.status.2"),value:2}]),le=E(()=>[{label:u("corp.dict.is_public.true"),value:"true"},{label:u("corp.dict.is_public.false"),value:"false"}]),D=async(t={...k})=>{i(!0);try{const{data:l}=await ct(t);T.value=l.items,I.current=t.current,I.pageSize=t.pageSize,I.total=l.paging.total}catch{}finally{i(!1)}},$=()=>{D({...k,...n.value})},oe=t=>{D({...k,...n.value,current:t})},ne=t=>{k.pageSize=t,D({...k,...n.value})};D();const ue=()=>{n.value=O()},se=async t=>{i(!0);try{await rt(t),y.$message.success("\u64CD\u4F5C\u6210\u529F"),$()}catch{}finally{i(!1)}},ie=async t=>{i(!0);try{await dt(t),y.$message.success("\u64CD\u4F5C\u6210\u529F"),$()}catch{}finally{i(!1)}},ce=(t,l)=>{z.value=t},re=(t,l,c)=>{t?b.value.splice(c,0,l):b.value=F.value.filter(r=>r.dataIndex!==l.dataIndex)},H=(t,l,c,r=!1)=>{const _=r?L(t):t;return l>-1&&c>-1&&_.splice(l,1,_.splice(c,1,_[l]).pop()),_},de=t=>{t&&Pe(()=>{const l=document.getElementById("tableSetting");new mt(l,{onEnd(c){const{oldIndex:r,newIndex:_}=c;H(b.value,r,_),H(F.value,r,_)}})})};Ue(()=>te.value,t=>{b.value=L(t),b.value.forEach((l,c)=>{l.checked=!0}),F.value=L(b.value)},{deep:!0,immediate:!0});const pe=t=>{h.value=t,C.value=!t.length},x=t=>{if(h.value.length===0)y.$message.info("\u8BF7\u9009\u62E9\u8981\u64CD\u4F5C\u7684\u6570\u636E");else{let l=`\u662F\u5426\u786E\u5B9A\u64CD\u4F5C\u6240\u9009\u7684${h.value.length}\u6761\u6570\u636E?`;switch(t.action){case"status":t.value===1?l=`\u662F\u5426\u786E\u5B9A\u542F\u7528\u6240\u9009\u7684${h.value.length}\u6761\u6570\u636E?`:l=`\u662F\u5426\u786E\u5B9A\u7981\u7528\u6240\u9009\u7684${h.value.length}\u6761\u6570\u636E?`;break;case"delete":l=`\u662F\u5426\u786E\u5B9A\u5220\u9664\u6240\u9009\u7684${h.value.length}\u6761\u6570\u636E?`;break}y.$modal.warning({title:"\u8B66\u544A",titleAlign:"center",content:l,hideCancel:!1,onOk:()=>{i(!0),t.ids=h.value,pt(t).then(c=>{i(!1),y.$message.success("\u64CD\u4F5C\u6210\u529F"),$(),R.value.selectAll(!1)})}})}};return(t,l)=>{const c=Ie,r=Le,_=Oe,U=Te,V=Re,d=He,M=Me,me=je,A=qe,_e=Ge,j=Je,fe=Ke,g=Qe,q=De,G=We,N=Xe,ge=Ee,ve=Ye,be=Ze,he=Se,Ce=Be,ye=et,ke=tt,J=at,$e=lt,Ve=ot,we=nt;return w(),S("div",_t,[e(_,{class:"container-breadcrumb"},{default:a(()=>[e(r,null,{default:a(()=>[e(c)]),_:1}),e(r,null,{default:a(()=>[s(m(t.$t("menu.corp")),1)]),_:1}),e(r,null,{default:a(()=>[s(m(t.$t("menu.corp.list")),1)]),_:1})]),_:1}),e(we,{class:"general-card",bordered:!1,"header-style":{padding:"20px"},"body-style":{padding:"25px 20px 20px 20px"}},{default:a(()=>[e(A,null,{default:a(()=>[e(d,{flex:1},{default:a(()=>[e(_e,{model:n.value,"label-col-props":{span:5},"wrapper-col-props":{span:18},"label-align":"left"},{default:a(()=>[e(A,{gutter:16},{default:a(()=>[e(d,{span:8},{default:a(()=>[e(V,{field:"name",label:t.$t("corp.form.name")},{default:a(()=>[e(U,{modelValue:n.value.name,"onUpdate:modelValue":l[0]||(l[0]=o=>n.value.name=o),placeholder:t.$t("corp.form.name.placeholder"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1}),e(d,{span:8},{default:a(()=>[e(V,{field:"code",label:t.$t("corp.form.code")},{default:a(()=>[e(U,{modelValue:n.value.code,"onUpdate:modelValue":l[1]||(l[1]=o=>n.value.code=o),placeholder:t.$t("corp.form.code.placeholder"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1}),e(d,{span:8},{default:a(()=>[e(V,{field:"remark",label:t.$t("corp.form.remark")},{default:a(()=>[e(U,{modelValue:n.value.remark,"onUpdate:modelValue":l[2]||(l[2]=o=>n.value.remark=o),placeholder:t.$t("corp.form.remark.placeholder"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1}),e(d,{span:8},{default:a(()=>[e(V,{field:"is_public",label:t.$t("corp.form.is_public")},{default:a(()=>[e(M,{modelValue:n.value.is_public,"onUpdate:modelValue":l[3]||(l[3]=o=>n.value.is_public=o),options:B(le),placeholder:t.$t("corp.form.selectDefault"),"allow-clear":""},null,8,["modelValue","options","placeholder"])]),_:1},8,["label"])]),_:1}),e(d,{span:8},{default:a(()=>[e(V,{field:"status",label:t.$t("corp.form.status")},{default:a(()=>[e(M,{modelValue:n.value.status,"onUpdate:modelValue":l[4]||(l[4]=o=>n.value.status=o),options:B(ae),placeholder:t.$t("corp.form.selectDefault"),"allow-clear":""},null,8,["modelValue","options","placeholder"])]),_:1},8,["label"])]),_:1}),e(d,{span:8},{default:a(()=>[e(V,{field:"updated_at",label:t.$t("corp.form.updated_at")},{default:a(()=>[e(me,{modelValue:n.value.updated_at,"onUpdate:modelValue":l[5]||(l[5]=o=>n.value.updated_at=o),style:{width:"100%"}},null,8,["modelValue"])]),_:1},8,["label"])]),_:1})]),_:1})]),_:1},8,["model"])]),_:1}),e(j,{style:{height:"84px"},direction:"vertical"}),e(d,{flex:"86px",style:{"text-align":"right"}},{default:a(()=>[e(G,{direction:"vertical",size:18},{default:a(()=>[e(g,{type:"primary",onClick:$},{icon:a(()=>[e(fe)]),default:a(()=>[s(" "+m(t.$t("corp.form.search")),1)]),_:1}),e(g,{onClick:ue},{icon:a(()=>[e(q)]),default:a(()=>[s(" "+m(t.$t("corp.form.reset")),1)]),_:1})]),_:1})]),_:1})]),_:1}),e(j,{style:{"margin-top":"0","margin-bottom":"16px"}}),e(A,{style:{"margin-bottom":"16px"}},{default:a(()=>[e(d,{span:12},{default:a(()=>[e(G,null,{default:a(()=>[e(g,{type:"primary",onClick:l[6]||(l[6]=o=>t.$router.push({name:"CorpCreate"}))},{default:a(()=>[s(m(t.$t("corp.operation.create")),1)]),_:1}),e(g,{type:"primary",status:"success",disabled:C.value,title:C.value?"\u8BF7\u9009\u62E9\u8981\u64CD\u4F5C\u7684\u6570\u636E":"",onClick:l[7]||(l[7]=o=>x({action:"status",value:1}))},{default:a(()=>[s(" \u542F\u7528 ")]),_:1},8,["disabled","title"]),e(g,{type:"primary",status:"danger",disabled:C.value,title:C.value?"\u8BF7\u9009\u62E9\u8981\u64CD\u4F5C\u7684\u6570\u636E":"",onClick:l[8]||(l[8]=o=>x({action:"status",value:2}))},{default:a(()=>[s(" \u7981\u7528 ")]),_:1},8,["disabled","title"]),e(g,{type:"primary",status:"danger",disabled:C.value,title:C.value?"\u8BF7\u9009\u62E9\u8981\u64CD\u4F5C\u7684\u6570\u636E":"",onClick:l[9]||(l[9]=o=>x({action:"delete"}))},{default:a(()=>[s(" \u5220\u9664 ")]),_:1},8,["disabled","title"])]),_:1})]),_:1}),e(d,{span:12,style:{display:"flex",height:"32px","align-items":"center","justify-content":"end"}},{default:a(()=>[e(N,{content:t.$t("actions.refresh")},{default:a(()=>[v("div",{class:"action-icon",onClick:$},[e(q,{size:"18"})])]),_:1},8,["content"]),e(be,{onSelect:ce},{content:a(()=>[(w(!0),S(Q,null,W(B(ee),o=>(w(),Ae(ve,{key:o.value,value:o.value,class:Ne({active:o.value===z.value})},{default:a(()=>[v("span",null,m(o.name),1)]),_:2},1032,["value","class"]))),128))]),default:a(()=>[e(N,{content:t.$t("actions.density")},{default:a(()=>[v("div",ft,[e(ge,{size:"18"})])]),_:1},8,["content"])]),_:1}),e(N,{content:t.$t("actions.column_setting")},{default:a(()=>[e(ke,{trigger:"click",position:"bl",onPopupVisibleChange:de},{content:a(()=>[v("div",vt,[(w(!0),S(Q,null,W(F.value,(o,f)=>(w(),S("div",{key:o.dataIndex,class:"setting"},[v("div",bt,[e(Ce)]),v("div",null,[e(ye,{modelValue:o.checked,"onUpdate:modelValue":P=>o.checked=P,onChange:P=>re(P,o,f)},null,8,["modelValue","onUpdate:modelValue","onChange"])]),v("div",ht,m(o.title==="#"?"\u5E8F\u5217\u53F7":o.title),1)]))),128))])]),default:a(()=>[v("div",gt,[e(he,{size:"18"})])]),_:1})]),_:1},8,["content"])]),_:1})]),_:1}),e(Ve,{ref_key:"tableRef",ref:R,"row-key":"id",loading:B(Z),pagination:I,columns:b.value,data:T.value,bordered:!1,size:z.value,"row-selection":X,onPageChange:oe,onPageSizeChange:ne,onSelectionChange:pe},{is_public:a(({record:o})=>[e(J,{modelValue:o.is_public,"onUpdate:modelValue":f=>o.is_public=f,"checked-value":!0,"unchecked-value":!1,onChange:f=>se({id:`${o.id}`,is_public:`${o.is_public}`})},null,8,["modelValue","onUpdate:modelValue","onChange"])]),remark:a(({record:o})=>[s(m(o.remark||"-"),1)]),status:a(({record:o})=>[e(J,{modelValue:o.status,"onUpdate:modelValue":f=>o.status=f,"checked-value":1,"unchecked-value":2,onChange:f=>ie({id:`${o.id}`,status:Number(`${o.status}`)})},null,8,["modelValue","onUpdate:modelValue","onChange"])]),operations:a(({record:o})=>[e(g,{type:"text",size:"small",onClick:f=>t.$router.push({name:"CorpUpdate",query:{id:`${o.id}`}})},{default:a(()=>[s(m(t.$t("corp.columns.operations.update")),1)]),_:2},1032,["onClick"]),e($e,{content:"\u4F60\u786E\u5B9A\u8981\u5220\u9664\u5417?",onOk:f=>Y({id:`${o.id}`})},{default:a(()=>[e(g,{type:"text",size:"small"},{default:a(()=>[s(m(t.$t("corp.columns.operations.delete")),1)]),_:1})]),_:2},1032,["onOk"])]),_:1},8,["loading","pagination","columns","data","size","row-selection"])]),_:1})])}}});const Xt=ze(yt,[["__scopeId","data-v-652cb721"]]);export{Xt as default};
