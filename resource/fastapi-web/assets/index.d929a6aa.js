import{u as ae,K as Ee,L as Be,b as M,M as Ae,N as Ne,O as Le,i as Ue,p as Pe,y as qe,z as Oe,_ as Te}from"./index.e0d6ab41.js";/* empty css                *//* empty css                *//* empty css               *//* empty css               *//* empty css              *//* empty css                *//* empty css               *//* empty css               *//* empty css                *//* empty css                *//* empty css               *//* empty css               *//* empty css                *//* empty css                *//* empty css               *//* empty css                *//* empty css                *//* empty css               *//* empty css               */import{d as le,e as v,B as a,C as p,aH as e,aG as t,u as l,aD as _,aM as s,aL as b,bJ as je,bK as He,bL as Me,bM as Re,bN as Ge,r as Z,c as R,w as Je,F as S,aJ as ee,aI as te,D as Ke,n as Qe,aK as We,aF as Xe,b1 as Ye,b2 as Ze,bC as et,bB as tt,bD as at,b5 as lt,bE as ot,ab as nt,aU as st,bi as it,bj as ut,bl as dt,bm as rt,b4 as ct,bF as pt,aT as _t,bG as mt,bH as ft,aV as gt,bI as bt,g as vt}from"./arco.91d8d802.js";import{u as oe}from"./loading.d8a03711.js";import{c as G,S as yt}from"./sortable.esm.2649578f.js";/* empty css                *//* empty css                */import"./chart.1c4d013e.js";import"./vue.90059513.js";const ht={style:{margin:"10px 0 30px 10px"}},kt={key:1},wt={key:1},Ct={key:1,style:{"max-height":"110px",display:"block",overflow:"auto"}},xt={key:1,style:{"max-height":"110px",display:"block",overflow:"auto"}},$t={key:1,style:{"max-height":"110px",display:"block",overflow:"auto"}},St={key:1,style:{"max-height":"110px",display:"block",overflow:"auto"}},Ft={key:1,style:{"max-height":"110px",display:"block",overflow:"auto"}},Dt={key:1,style:{"max-height":"110px",display:"block",overflow:"auto"}},It={key:1,style:{"max-height":"110px",display:"block",overflow:"auto"}},Vt={key:1,style:{"max-height":"110px",display:"block",overflow:"auto"}},zt={key:1},Et={key:1},Bt={key:1,style:{"max-height":"110px",display:"block",overflow:"auto"}},At={key:1},Nt={key:1},Lt={key:1,style:{"max-height":"110px",display:"block",overflow:"auto"}},Ut={key:1},Pt={key:1},qt={key:1},Ot={key:1},Tt={key:1},jt={key:1,style:{"max-height":"110px",display:"block",overflow:"auto"}},Ht={key:1},Mt={key:1},Rt={key:1},Gt={name:"SiteConfigDetail"},Jt=le({...Gt,props:{id:{type:String,default:""}},setup(J){const F=J,{t:m}=ae(),{loading:f,setLoading:N}=oe(!0),r=v({});return(async(g={id:F.id})=>{N(!0);try{const{data:V}=await Ee(g);r.value=V}catch{}finally{N(!1)}})(),(g,V)=>{const i=je,d=He,c=Me,z=Re,C=Ge;return a(),p("div",ht,[e(C,{column:2,bordered:"","value-style":{width:"350px",padding:"5px 8px 5px 20px"}},{default:t(()=>[e(c,{label:l(m)("site.config.detail.domain"),span:2},{default:t(()=>[l(f)?(a(),_(d,{key:0,animation:!0},{default:t(()=>[e(i,{rows:1})]),_:1})):(a(),p("span",kt,s(r.value.domain),1))]),_:1},8,["label"]),e(c,{label:l(m)("site.config.detail.title")},{default:t(()=>[l(f)?(a(),_(d,{key:0,animation:!0},{default:t(()=>[e(i,{rows:1})]),_:1})):(a(),p("span",wt,s(r.value.title),1))]),_:1},8,["label"]),e(c,{label:l(m)("site.config.detail.logo"),span:2},{default:t(()=>[l(f)?(a(),_(d,{key:0,animation:!0},{default:t(()=>[e(i,{rows:1})]),_:1})):(a(),p("span",Ct,s(r.value.logo),1))]),_:1},8,["label"]),e(c,{label:l(m)("site.config.detail.favicon")},{default:t(()=>[l(f)?(a(),_(d,{key:0,animation:!0},{default:t(()=>[e(i,{rows:1})]),_:1})):(a(),p("span",xt,s(r.value.favicon),1))]),_:1},8,["label"]),e(c,{label:l(m)("site.config.detail.avatar"),span:2},{default:t(()=>[l(f)?(a(),_(d,{key:0,animation:!0},{default:t(()=>[e(i,{rows:1})]),_:1})):(a(),p("span",$t,s(r.value.avatar||"-"),1))]),_:1},8,["label"]),e(c,{label:l(m)("site.config.detail.bg_img")},{default:t(()=>[l(f)?(a(),_(d,{key:0,animation:!0},{default:t(()=>[e(i,{rows:1})]),_:1})):(a(),p("span",St,s(r.value.bg_img||"-"),1))]),_:1},8,["label"]),e(c,{label:l(m)("site.config.detail.copyright"),span:2},{default:t(()=>[l(f)?(a(),_(d,{key:0,animation:!0},{default:t(()=>[e(i,{rows:1})]),_:1})):(a(),p("span",Ft,s(r.value.copyright||"-"),1))]),_:1},8,["label"]),e(c,{label:l(m)("site.config.detail.jump_url")},{default:t(()=>[l(f)?(a(),_(d,{key:0,animation:!0},{default:t(()=>[e(i,{rows:1})]),_:1})):(a(),p("span",Dt,s(r.value.jump_url||"-"),1))]),_:1},8,["label"]),e(c,{label:l(m)("site.config.detail.keywords"),span:2},{default:t(()=>[l(f)?(a(),_(d,{key:0,animation:!0},{default:t(()=>[e(i,{rows:1})]),_:1})):(a(),p("span",It,s(r.value.keywords||"-"),1))]),_:1},8,["label"]),e(c,{label:l(m)("site.config.detail.description")},{default:t(()=>[l(f)?(a(),_(d,{key:0,animation:!0},{default:t(()=>[e(i,{rows:1})]),_:1})):(a(),p("span",Vt,s(r.value.description||"-"),1))]),_:1},8,["label"]),e(c,{label:l(m)("site.config.detail.icp_beian"),span:2},{default:t(()=>[l(f)?(a(),_(d,{key:0,animation:!0},{default:t(()=>[e(i,{rows:1})]),_:1})):(a(),p("span",zt,s(r.value.icp_beian||"-"),1))]),_:1},8,["label"]),e(c,{label:l(m)("site.config.detail.ga_beian")},{default:t(()=>[l(f)?(a(),_(d,{key:0,animation:!0},{default:t(()=>[e(i,{rows:1})]),_:1})):(a(),p("span",Et,s(r.value.ga_beian||"-"),1))]),_:1},8,["label"]),e(c,{label:l(m)("site.config.detail.register_tips"),span:2},{default:t(()=>[l(f)?(a(),_(d,{key:0,animation:!0},{default:t(()=>[e(i,{rows:1})]),_:1})):(a(),p("span",Bt,s(r.value.register_tips||"-"),1))]),_:1},8,["label"]),e(c,{label:l(m)("site.config.detail.grant_quota")},{default:t(()=>[l(f)?(a(),_(d,{key:0,animation:!0},{default:t(()=>[e(i,{rows:1})]),_:1})):(a(),p("span",At,s(r.value.grant_quota||"-"),1))]),_:1},8,["label"]),e(c,{label:l(m)("site.config.detail.quota_expires_at"),span:2},{default:t(()=>[l(f)?(a(),_(d,{key:0,animation:!0},{default:t(()=>[e(i,{rows:1})]),_:1})):(a(),p("span",Nt,s(r.value.quota_expires_at||"-"),1))]),_:1},8,["label"]),e(c,{label:l(m)("site.config.detail.support_email_suffix")},{default:t(()=>{var h,B;return[l(f)?(a(),_(d,{key:0,animation:!0},{default:t(()=>[e(i,{rows:1})]),_:1})):(a(),p("span",Lt,s(((B=(h=r.value)==null?void 0:h.support_email_suffix)==null?void 0:B.join(`
`))||"-"),1))]}),_:1},8,["label"]),e(c,{label:l(m)("site.config.detail.host"),span:2},{default:t(()=>[l(f)?(a(),_(d,{key:0,animation:!0},{default:t(()=>[e(i,{rows:1})]),_:1})):(a(),p("span",Ut,s(r.value.host||"-"),1))]),_:1},8,["label"]),e(c,{label:l(m)("site.config.detail.port")},{default:t(()=>[l(f)?(a(),_(d,{key:0,animation:!0},{default:t(()=>[e(i,{rows:1})]),_:1})):(a(),p("span",Pt,s(r.value.port||"-"),1))]),_:1},8,["label"]),e(c,{label:l(m)("site.config.detail.user_name"),span:2},{default:t(()=>[l(f)?(a(),_(d,{key:0,animation:!0},{default:t(()=>[e(i,{rows:1})]),_:1})):(a(),p("span",qt,s(r.value.user_name||"-"),1))]),_:1},8,["label"]),e(c,{label:l(m)("site.config.detail.password")},{default:t(()=>[l(f)?(a(),_(d,{key:0,animation:!0},{default:t(()=>[e(i,{rows:1})]),_:1})):(a(),p("span",Ot,s(r.value.password||"-"),1))]),_:1},8,["label"]),e(c,{label:l(m)("site.config.detail.from_name"),span:2},{default:t(()=>[l(f)?(a(),_(d,{key:0,animation:!0},{default:t(()=>[e(i,{rows:1})]),_:1})):(a(),p("span",Tt,s(r.value.from_name||"-"),1))]),_:1},8,["label"]),e(c,{label:l(m)("common.remark")},{default:t(()=>[l(f)?(a(),_(d,{key:0,animation:!0},{default:t(()=>[e(i,{rows:1})]),_:1})):(a(),p("span",jt,s(r.value.remark||"-"),1))]),_:1},8,["label"]),e(c,{label:l(m)("common.status"),span:2},{default:t(()=>[l(f)?(a(),_(d,{key:0,animation:!0},{default:t(()=>[e(i,{rows:1})]),_:1})):(a(),p("span",Ht,[r.value.status===1?(a(),_(z,{key:0,color:"green"},{default:t(()=>[b(s(g.$t(`dict.status.${r.value.status}`)),1)]),_:1})):(a(),_(z,{key:1,color:"red"},{default:t(()=>[b(s(g.$t(`dict.status.${r.value.status}`)),1)]),_:1}))]))]),_:1},8,["label"]),e(c,{label:l(m)("common.created_at")},{default:t(()=>[l(f)?(a(),_(d,{key:0,animation:!0},{default:t(()=>[e(i,{rows:1})]),_:1})):(a(),p("span",Mt,s(r.value.created_at),1))]),_:1},8,["label"]),e(c,{label:l(m)("common.updated_at"),span:2},{default:t(()=>[l(f)?(a(),_(d,{key:0,animation:!0},{default:t(()=>[e(i,{rows:1})]),_:1})):(a(),p("span",Rt,s(r.value.updated_at),1))]),_:1},8,["label"])]),_:1})])}}}),Kt={class:"container"},Qt={class:"action-icon"},Wt={class:"action-icon"},Xt={id:"tableSetting"},Yt={style:{"margin-right":"4px",cursor:"move"}},Zt={class:"title"},ea={name:"SiteConfigList"},ta=le({...ea,setup(J){const{proxy:F}=vt(),m=Z({type:"checkbox",showCheckedAll:!0,onlyCurrent:!1}),f=async o=>{w(!0);try{await Be(o),F.$message.success("\u5220\u9664\u6210\u529F"),M().init(),A()}catch{}finally{w(!1)}},N=()=>({user_id:v(),domain:"",title:"",register_tips:"",logo:"",status:v()}),{loading:r,setLoading:w}=oe(!0),{t:g}=ae(),V=v([]),i=v(N()),d=v([]),c=v([]),z=v("medium"),C=v([]),h=v(!0),B=v(),E={current:1,pageSize:20,showTotal:!0,showPageSize:!0,pageSizeOptions:[20,50,100,500,1e3]},L=Z({...E}),ne=R(()=>[{name:g("size.mini"),value:"mini"},{name:g("size.small"),value:"small"},{name:g("size.medium"),value:"medium"},{name:g("size.large"),value:"large"}]),se=R(()=>[{title:g("site.config.columns.domain"),dataIndex:"domain",slotName:"domain",align:"center",ellipsis:!0,tooltip:!0},{title:g("site.config.columns.title"),dataIndex:"title",slotName:"title",align:"center",ellipsis:!0,tooltip:!0},{title:g("site.config.columns.register_tips"),dataIndex:"register_tips",slotName:"register_tips",align:"center",ellipsis:!0,tooltip:!0},{title:g("common.remark"),dataIndex:"remark",slotName:"remark",align:"center",ellipsis:!0,tooltip:!0},{title:g("common.status"),dataIndex:"status",slotName:"status",align:"center",width:65},{title:g("common.updated_at"),dataIndex:"updated_at",slotName:"updated_at",align:"center",width:132},{title:g("common.operations"),dataIndex:"operations",slotName:"operations",align:"center",width:170}]),ie=R(()=>[{label:g("dict.status.1"),value:1},{label:g("dict.status.2"),value:2}]),U=async(o={...E})=>{w(!0);try{const{data:n}=await Ae(o);V.value=n.items,L.current=o.current,L.pageSize=o.pageSize,L.total=n.paging.total}catch{}finally{w(!1)}},A=()=>{U({...E,...i.value})},ue=o=>{U({...E,...i.value,current:o})},de=o=>{E.pageSize=o,U({...E,...i.value})};U();const re=async o=>{w(!0);try{await Ne(o),F.$message.success("\u64CD\u4F5C\u6210\u529F"),M().init(),A()}catch{}finally{w(!1)}},ce=(o,n)=>{z.value=o},pe=(o,n,y)=>{o?d.value.splice(y,0,n):d.value=c.value.filter(k=>k.dataIndex!==n.dataIndex)},K=(o,n,y,k=!1)=>{const x=k?G(o):o;return n>-1&&y>-1&&x.splice(n,1,x.splice(y,1,x[n]).pop()),x},_e=o=>{o&&Qe(()=>{const n=document.getElementById("tableSetting");new yt(n,{onEnd(y){const{oldIndex:k,newIndex:x}=y;K(d.value,k,x),K(c.value,k,x)}})})};Je(()=>se.value,o=>{d.value=G(o),d.value.forEach((n,y)=>{n.checked=!0}),c.value=G(d.value)},{deep:!0,immediate:!0});const me=o=>{C.value=o,h.value=!o.length},P=o=>{if(C.value.length===0)F.$message.info("\u8BF7\u9009\u62E9\u8981\u64CD\u4F5C\u7684\u6570\u636E");else{let n=`\u662F\u5426\u786E\u5B9A\u64CD\u4F5C\u6240\u9009\u7684${C.value.length}\u6761\u6570\u636E?`;switch(o.action){case"status":o.value===1?n=`\u662F\u5426\u786E\u5B9A\u542F\u7528\u6240\u9009\u7684${C.value.length}\u6761\u6570\u636E?`:n=`\u662F\u5426\u786E\u5B9A\u7981\u7528\u6240\u9009\u7684${C.value.length}\u6761\u6570\u636E?`;break;case"delete":n=`\u662F\u5426\u786E\u5B9A\u5220\u9664\u6240\u9009\u7684${C.value.length}\u6761\u6570\u636E?`;break}F.$modal.warning({title:"\u8B66\u544A",titleAlign:"start",content:n,hideCancel:!1,onOk:()=>{w(!0),o.ids=C.value,Le(o).then(y=>{w(!1),F.$message.success("\u64CD\u4F5C\u6210\u529F"),M().init(),A(),B.value.selectAll(!1)})}})}},q=v(!1),Q=v(),fe=o=>{q.value=!0,Q.value=o},ge=()=>{q.value=!1};return(o,n)=>{const y=Ue,k=We,x=Xe,W=Ye,O=Ze,D=et,be=tt,T=at,ve=lt,X=ot,ye=nt,$=st,Y=it,he=Pe,j=ut,ke=qe,we=dt,Ce=rt,xe=Oe,$e=ct,Se=pt,Fe=_t,De=mt,Ie=ft,Ve=gt,ze=bt;return a(),p("div",Kt,[e(x,{class:"container-breadcrumb"},{default:t(()=>[e(k,null,{default:t(()=>[e(y)]),_:1}),e(k,null,{default:t(()=>[b(s(o.$t("menu.sys")),1)]),_:1}),e(k,null,{default:t(()=>[b(s(o.$t("menu.site.config")),1)]),_:1})]),_:1}),e(ze,{class:"general-card",bordered:!1,"header-style":{padding:"20px"},"body-style":{padding:"25px 20px 20px 20px"}},{default:t(()=>[e(T,null,{default:t(()=>[e(D,{flex:1},{default:t(()=>[e(ve,{model:i.value,"label-col-props":{span:5},"wrapper-col-props":{span:18},"label-align":"left"},{default:t(()=>[e(T,{gutter:16},{default:t(()=>[e(D,{span:8},{default:t(()=>[e(O,{field:"domain",label:o.$t("site.config.form.domain")},{default:t(()=>[e(W,{modelValue:i.value.domain,"onUpdate:modelValue":n[0]||(n[0]=u=>i.value.domain=u),placeholder:o.$t("site.config.form.domain.placeholder"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1}),e(D,{span:8},{default:t(()=>[e(O,{field:"title",label:o.$t("site.config.form.title")},{default:t(()=>[e(W,{modelValue:i.value.title,"onUpdate:modelValue":n[1]||(n[1]=u=>i.value.title=u),placeholder:o.$t("site.config.form.title.placeholder"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1}),e(D,{span:8},{default:t(()=>[e(O,{field:"status",label:o.$t("common.status")},{default:t(()=>[e(be,{modelValue:i.value.status,"onUpdate:modelValue":n[2]||(n[2]=u=>i.value.status=u),options:l(ie),placeholder:o.$t("common.all"),"allow-clear":""},null,8,["modelValue","options","placeholder"])]),_:1},8,["label"])]),_:1})]),_:1})]),_:1},8,["model"])]),_:1}),e(X,{style:{height:"42px"},direction:"vertical"}),e(D,{flex:"86px",style:{"text-align":"right"}},{default:t(()=>[e(Y,{direction:"vertical",size:18},{default:t(()=>[e($,{type:"primary",onClick:A},{icon:t(()=>[e(ye)]),default:t(()=>[b(" "+s(o.$t("button.search")),1)]),_:1})]),_:1})]),_:1})]),_:1}),e(X,{style:{"margin-top":"0","margin-bottom":"16px"}}),e(T,{style:{"margin-bottom":"16px"}},{default:t(()=>[e(D,{span:12},{default:t(()=>[e(Y,null,{default:t(()=>[e($,{type:"primary",onClick:n[3]||(n[3]=u=>o.$router.push({name:"SiteConfigCreate"}))},{default:t(()=>[b(s(o.$t("button.create")),1)]),_:1}),e($,{type:"primary",status:"success",disabled:h.value,title:h.value?"\u8BF7\u9009\u62E9\u8981\u64CD\u4F5C\u7684\u6570\u636E":"",onClick:n[4]||(n[4]=u=>P({action:"status",value:1}))},{default:t(()=>[b(" \u542F\u7528 ")]),_:1},8,["disabled","title"]),e($,{type:"primary",status:"danger",disabled:h.value,title:h.value?"\u8BF7\u9009\u62E9\u8981\u64CD\u4F5C\u7684\u6570\u636E":"",onClick:n[5]||(n[5]=u=>P({action:"status",value:2}))},{default:t(()=>[b(" \u7981\u7528 ")]),_:1},8,["disabled","title"]),e($,{type:"primary",status:"danger",disabled:h.value,title:h.value?"\u8BF7\u9009\u62E9\u8981\u64CD\u4F5C\u7684\u6570\u636E":"",onClick:n[6]||(n[6]=u=>P({action:"delete"}))},{default:t(()=>[b(" \u5220\u9664 ")]),_:1},8,["disabled","title"])]),_:1})]),_:1}),e(D,{span:12,style:{display:"flex",height:"32px","align-items":"center","justify-content":"end"}},{default:t(()=>[e(j,{content:o.$t("actions.refresh")},{default:t(()=>[S("div",{class:"action-icon",onClick:A},[e(he,{size:"18"})])]),_:1},8,["content"]),e(Ce,{onSelect:ce},{content:t(()=>[(a(!0),p(ee,null,te(l(ne),u=>(a(),_(we,{key:u.value,value:u.value,class:Ke({active:u.value===z.value})},{default:t(()=>[S("span",null,s(u.name),1)]),_:2},1032,["value","class"]))),128))]),default:t(()=>[e(j,{content:o.$t("actions.density")},{default:t(()=>[S("div",Qt,[e(ke,{size:"18"})])]),_:1},8,["content"])]),_:1}),e(j,{content:o.$t("actions.column_setting")},{default:t(()=>[e(Se,{trigger:"click",position:"bl",onPopupVisibleChange:_e},{content:t(()=>[S("div",Xt,[(a(!0),p(ee,null,te(c.value,(u,I)=>(a(),p("div",{key:u.dataIndex,class:"setting"},[S("div",Yt,[e(xe)]),S("div",null,[e($e,{modelValue:u.checked,"onUpdate:modelValue":H=>u.checked=H,onChange:H=>pe(H,u,I)},null,8,["modelValue","onUpdate:modelValue","onChange"])]),S("div",Zt,s(u.title==="#"?"\u5E8F\u5217\u53F7":u.title),1)]))),128))])]),default:t(()=>[S("div",Wt,[e(y,{size:"18"})])]),_:1})]),_:1},8,["content"])]),_:1})]),_:1}),e(Ie,{ref_key:"tableRef",ref:B,"row-key":"id",loading:l(r),pagination:L,columns:d.value,data:V.value,bordered:!1,size:z.value,"row-selection":m,onPageChange:ue,onPageSizeChange:de,onSelectionChange:me},{register_tips:t(({record:u})=>[b(s(u.register_tips||"-"),1)]),remark:t(({record:u})=>[b(s(u.remark||"-"),1)]),status:t(({record:u})=>[e(Fe,{modelValue:u.status,"onUpdate:modelValue":I=>u.status=I,"checked-value":1,"unchecked-value":2,onChange:I=>re({id:`${u.id}`,status:Number(`${u.status}`)})},null,8,["modelValue","onUpdate:modelValue","onChange"])]),operations:t(({record:u})=>[e($,{type:"text",size:"small",onClick:I=>fe(u.id)},{default:t(()=>[b(s(o.$t("operations.view")),1)]),_:2},1032,["onClick"]),e($,{type:"text",size:"small",onClick:I=>o.$router.push({name:"SiteConfigUpdate",query:{id:`${u.id}`}})},{default:t(()=>[b(s(o.$t("operations.update")),1)]),_:2},1032,["onClick"]),e(De,{content:"\u4F60\u786E\u5B9A\u8981\u5220\u9664\u5417?",onOk:I=>f({id:`${u.id}`})},{default:t(()=>[e($,{type:"text",size:"small"},{default:t(()=>[b(s(o.$t("operations.delete")),1)]),_:1})]),_:2},1032,["onOk"])]),_:1},8,["loading","pagination","columns","data","size","row-selection"]),e(Ve,{title:o.$t("menu.site.config.detail"),"unmount-on-close":"","render-to-body":"",width:700,footer:!1,visible:q.value,onCancel:ge},{default:t(()=>[e(Jt,{id:Q.value},null,8,["id"])]),_:1},8,["title","visible"])]),_:1})])}}});const Ia=Te(ta,[["__scopeId","data-v-c71a07b3"]]);export{Ia as default};
