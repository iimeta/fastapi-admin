import{u as dt,F as rt,p as _t,y as mt,i as ct,z as pt,_ as ft}from"./index.387a2346.js";/* empty css               *//* empty css               *//* empty css                *//* empty css                *//* empty css               *//* empty css              *//* empty css               *//* empty css               *//* empty css                *//* empty css                *//* empty css               *//* empty css               */import{c as ie,S as vt}from"./sortable.esm.f8f875e8.js";/* empty css               *//* empty css                *//* empty css                *//* empty css               *//* empty css               *//* empty css              *//* empty css              */import{d as yt,r as qe,e as h,c as de,w as Ie,bT as kt,B as t,C as _,aH as a,aG as e,aL as i,aM as n,bu as k,aD as o,aJ as G,aI as K,u as m,F as $,D as ht,aE as bt,bP as Ve,n as gt,g as wt,aK as Ft,aF as Et,bA as Dt,bB as Ct,b2 as $t,bC as Bt,b1 as At,aS as qt,bD as It,bE as Vt,b5 as xt,bF as St,ab as zt,aU as Lt,bi as Nt,bj as Pt,bl as Tt,bm as Ut,b4 as Mt,bG as Ot,b$ as Rt,bI as Yt,bM as Ht,bN as jt,ad as Jt,c0 as Gt,bO as Kt,ba as Qt,b6 as Wt,aV as Xt,a_ as Zt,bJ as ea}from"./arco.17b1a46f.js";import{u as ta}from"./loading.44762de3.js";import{q as aa}from"./common.ac936b7b.js";import{V as xe,q as la,a as oa,s as ua,b as na}from"./styles.3a020a6b.js";import{g as sa}from"./app.7b606d42.js";import{b as ia}from"./dashboard.b99d46b2.js";import{q as da}from"./model.c8c8863e.js";import{u as ra}from"./vue.32c094a4.js";import"./chart.d5ce7f1f.js";import"./base.87fcf6e2.js";const _a={class:"container"},ma={class:"action-icon"},ca={class:"action-icon"},pa={id:"tableSetting"},fa={style:{"margin-right":"4px",cursor:"move"}},va={class:"title"},ya={style:{margin:"10px 0 0 10px"}},ka={key:1},ha={key:1},ba={key:1},ga={key:1},wa={key:1},Fa={key:1},Ea={key:1},Da={key:1},Ca={key:1},$a={key:1},Ba={key:1},Aa={key:1},qa={key:1},Ia={key:1},Va={key:1},xa={key:1},Sa={key:1},za={key:1},La={key:1},Na={key:1},Pa={key:1},Ta={key:1},Ua={key:1},Ma={key:1},Oa={key:1},Ra={key:1},Ya={key:1},Ha={key:1},ja={key:1},Ja={key:1},Ga={key:1},Ka={key:1},Qa={key:1},Wa={key:1},Xa={key:1},Za={key:1},el={key:1},tl={key:1},al={key:1},ll={key:1},ol={key:1},ul={key:1},nl={key:1},sl={key:1},il={key:1},dl={key:1},rl={key:1},_l={key:1},ml={key:1},cl={key:1},pl={key:1},fl={key:1},vl={key:1},yl={key:1},kl={key:1},hl={key:1},bl={key:1},gl={key:1},wl={key:1},Fl={key:1},El={key:1},Dl={name:"ChatList"},Cl=yt({...Dl,setup($l){const Se=qe({type:"checkbox",showCheckedAll:!0,onlyCurrent:!1}),re=h([]);(async()=>{try{const{data:s}=await sa();re.value=s.items}catch{}})();const _e=h([]);(async()=>{try{const{data:s}=await da();_e.value=s.items}catch{}})();const me=()=>({app_id:h(),trace_id:h(),user_id:h(),key:"",models:[],total_time:h(),status:h(),req_time:[Ve().format("YYYY-MM-DD 00:00:00"),Ve().format("YYYY-MM-DD 23:59:59")]}),{loading:c,setLoading:I}=ta(!0),{t:b}=dt(),ce=h([]),g=h(me()),A=h([]),U=h([]),Q=h("medium"),q=h([]),ze=h(!0),W=h(),V={current:1,pageSize:10,showTotal:!0,showPageSize:!0,pageSizeOptions:[10,50,100,500,1e3]},M=qe({...V}),Le=de(()=>[{name:b("searchTable.size.mini"),value:"mini"},{name:b("searchTable.size.small"),value:"small"},{name:b("searchTable.size.medium"),value:"medium"},{name:b("searchTable.size.large"),value:"large"}]),x=localStorage.getItem("userRole"),pe=de(()=>[{title:b(x==="admin"?"chat.columns.user_id":"chat.columns.app_id"),dataIndex:x==="admin"?"user_id":"app_id",slotName:x==="admin"?"user_id":"app_id",align:"center",width:75},{title:b("chat.columns.model"),dataIndex:"model",slotName:"model",align:"center"},{title:b("chat.columns.prompt_tokens"),dataIndex:"prompt_tokens",slotName:"prompt_tokens",align:"center"},{title:b("chat.columns.completion_tokens"),dataIndex:"completion_tokens",slotName:"completion_tokens",align:"center"},{title:b("chat.columns.total_price"),dataIndex:"total_tokens",slotName:"total_tokens",align:"center"},{title:b("chat.columns.stream"),dataIndex:"stream",slotName:"stream",align:"center"},{title:b("chat.columns.conn_time"),dataIndex:"conn_time",slotName:"conn_time",align:"center"},{title:b("chat.columns.duration"),dataIndex:"duration",slotName:"duration",align:"center"},{title:b("chat.columns.total_time"),dataIndex:"total_time",slotName:"total_time",align:"center"},{title:b("chat.columns.internal_time"),dataIndex:"internal_time",slotName:"internal_time",align:"center"},{title:b("chat.columns.status"),dataIndex:"status",slotName:"status",align:"center",width:65},{title:b("chat.columns.req_time"),dataIndex:"req_time",slotName:"req_time",align:"center",width:132},{title:b("chat.columns.operations"),dataIndex:"operations",slotName:"operations",align:"center",width:75}]);x==="user"&&pe.value.splice(9,1);const fe=de(()=>[{label:b("chat.dict.status.1"),value:1},{label:b("chat.dict.status.2"),value:2},{label:b("chat.dict.status.-1"),value:-1}]);x==="admin"&&fe.value.push({label:b("chat.dict.status.3"),value:3},{label:b("chat.dict.status.-100"),value:-100});const O=async(s={...V,...g.value})=>{I(!0);try{const{data:r}=await la(s);ce.value=r.items,M.current=s.current,M.pageSize=s.pageSize,M.total=r.paging.total}catch{}finally{I(!1)}},X=()=>{O({...V,...g.value})},Ne=s=>{O({...V,...g.value,current:s})},Pe=s=>{V.pageSize=s,O({...V,...g.value})};O();const Te=()=>{g.value=me()},Ue=(s,r)=>{Q.value=s},Me=(s,r,w)=>{s?A.value.splice(w,0,r):A.value=U.value.filter(D=>D.dataIndex!==r.dataIndex)},ve=(s,r,w,D=!1)=>{const F=D?ie(s):s;return r>-1&&w>-1&&F.splice(r,1,F.splice(w,1,F[r]).pop()),F},Oe=s=>{s&&gt(()=>{const r=document.getElementById("tableSetting");new vt(r,{onEnd(w){const{oldIndex:D,newIndex:F}=w;ve(A.value,D,F),ve(U.value,D,F)}})})};Ie(()=>pe.value,s=>{A.value=ie(s),A.value.forEach((r,w)=>{r.checked=!0}),U.value=ie(A.value)},{deep:!0,immediate:!0});const Z=h(!1),{copy:Re,copied:ye}=ra(),{proxy:S}=wt(),l=h({}),Ye=async s=>{Z.value=!0,c.value=!0;try{const{data:r}=await oa({id:s});l.value=r}catch{}finally{c.value=!1}},He=()=>{Z.value=!1},z=s=>{Re(s)};Ie(ye,()=>{ye.value&&S.$message.success("\u590D\u5236\u6210\u529F")});const ke=h(0),he=h(0),R=async(s={...g.value})=>{const{data:r}=await ia(s);ke.value=r.rpm,he.value=r.tpm};R();let ee;ee=setInterval(()=>{R()},3e3),window.onblur=()=>{clearInterval(ee)},window.onfocus=()=>{R(),ee=setInterval(()=>{R()},3e3)};const be=h(),L=h(!1),Y=h({}),je=async s=>{var w;if(await((w=be.value)==null?void 0:w.validate())){L.value=!0,s(!1);return}s(),ge({req_time:Y.value.req_time})},Je=()=>{L.value=!1},ge=s=>{if(q.value.length===0&&!s.req_time){L.value=!0;return}I(!0),s.ids=q.value,ua(s).then(r=>{I(!1),S.$message.success("\u5BFC\u51FA\u6210\u529F"),W.value.selectAll(!1);const w=new Blob([r.data],{type:"application/vnd.ms-excel"}),D=window.URL.createObjectURL(w),F=document.createElement("a");F.href=D,F.setAttribute("download","\u804A\u5929\u65E5\u5FD7.xlsx"),document.body.appendChild(F),F.click(),document.body.removeChild(F),window.URL.revokeObjectURL(D)}).catch(r=>{S.$message.error("\u5BFC\u51FA\u5931\u8D25, \u8BF7\u8054\u7CFB\u7BA1\u7406\u5458",r)})},we=h(),N=h(!1),H=h({}),Ge=async s=>{var w;if(await((w=we.value)==null?void 0:w.validate())){N.value=!0,s(!1);return}s(),Fe({action:"time",value:H.value.value})},Ke=()=>{N.value=!1},Qe=s=>{q.value=s,ze.value=!s.length},Fe=s=>{if(q.value.length===0&&!s.value)N.value=!0;else{let r=`\u662F\u5426\u786E\u5B9A\u64CD\u4F5C\u6240\u9009\u7684${q.value.length}\u6761\u6570\u636E?`;switch(s.action){case"delete":r=`\u662F\u5426\u786E\u5B9A\u5220\u9664\u6240\u9009\u7684${q.value.length}\u6761\u6570\u636E?`;break;case"time":r=`\u662F\u5426\u786E\u5B9A\u5220\u9664${s.value[0]}\u81F3${s.value[1]}\u7684\u6570\u636E?`;break}S.$modal.warning({title:"\u8B66\u544A",titleAlign:"start",content:r,hideCancel:!1,onOk:()=>{I(!0),s.ids=q.value,na(s).then(w=>{I(!1),S.$message.success("\u64CD\u4F5C\u6210\u529F"),X(),W.value.selectAll(!1)})}})}};return(s,r)=>{const w=rt,D=Ft,F=Et,Ee=Dt,te=Ct,C=$t,E=Bt,ae=At,We=qt,le=It,oe=Vt,ue=xt,De=St,Xe=zt,P=Lt,Ce=_t,j=Nt,ne=Pt,Ze=mt,et=Tt,tt=Ut,at=ct,lt=pt,ot=Mt,ut=Ot,d=Rt,nt=Yt,p=Ht,f=jt,T=Jt,v=Gt,J=Kt,$e=Qt,Be=Wt,st=Xt,Ae=Zt,it=ea,y=kt("permission");return t(),_("div",_a,[a(F,{class:"container-breadcrumb"},{default:e(()=>[a(D,null,{default:e(()=>[a(w)]),_:1}),a(D,null,{default:e(()=>[i(n(s.$t("menu.chat")),1)]),_:1}),a(D,null,{default:e(()=>[i(n(s.$t("menu.chat.list")),1)]),_:1})]),_:1}),a(it,{class:"general-card",bordered:!1,"header-style":{padding:"20px"},"body-style":{padding:"25px 20px 20px 20px"}},{default:e(()=>[a(oe,null,{default:e(()=>[a(E,{flex:1},{default:e(()=>[a(ue,{model:g.value,"label-col-props":{span:5},"wrapper-col-props":{span:18},"label-align":"left"},{default:e(()=>[a(oe,{gutter:16},{default:e(()=>[k((t(),o(E,{span:8},{default:e(()=>[a(C,{field:"app_id",label:s.$t("chat.form.app_id")},{default:e(()=>[a(te,{modelValue:g.value.app_id,"onUpdate:modelValue":r[0]||(r[0]=u=>g.value.app_id=u),placeholder:s.$t("chat.form.selectDefault"),"allow-search":"","allow-clear":""},{default:e(()=>[(t(!0),_(G,null,K(re.value,u=>(t(),o(Ee,{key:u.app_id,value:u.app_id,label:u.name},null,8,["value","label"]))),128))]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1})),[[y,["user"]]]),k((t(),o(E,{span:8},{default:e(()=>[a(C,{field:"trace_id",label:s.$t("chat.form.trace_id")},{default:e(()=>[a(ae,{modelValue:g.value.trace_id,"onUpdate:modelValue":r[1]||(r[1]=u=>g.value.trace_id=u),placeholder:s.$t("chat.form.trace_id.placeholder"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1})),[[y,["admin"]]]),k((t(),o(E,{span:8},{default:e(()=>[a(C,{field:"user_id",label:s.$t("chat.form.user_id")},{default:e(()=>[a(ae,{modelValue:g.value.user_id,"onUpdate:modelValue":r[2]||(r[2]=u=>g.value.user_id=u),placeholder:s.$t("chat.form.user_id.placeholder"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1})),[[y,["admin"]]]),k((t(),o(E,{span:8},{default:e(()=>[a(C,{field:"key",label:s.$t("chat.form.key")},{default:e(()=>[a(ae,{modelValue:g.value.key,"onUpdate:modelValue":r[3]||(r[3]=u=>g.value.key=u),placeholder:s.$t("chat.form.key.placeholder"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1})),[[y,["user"]]]),a(E,{span:8},{default:e(()=>[a(C,{field:"models",label:s.$t("chat.form.models")},{default:e(()=>[a(te,{modelValue:g.value.models,"onUpdate:modelValue":r[4]||(r[4]=u=>g.value.models=u),placeholder:s.$t("chat.form.selectDefault"),"max-tag-count":2,multiple:"","allow-search":"","allow-clear":""},{default:e(()=>[(t(!0),_(G,null,K(_e.value,u=>(t(),o(Ee,{key:u.id,value:u.id,label:u.name},null,8,["value","label"]))),128))]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1}),a(E,{span:8},{default:e(()=>[a(C,{field:"total_time",label:s.$t("chat.form.total_time")},{default:e(()=>[a(We,{modelValue:g.value.total_time,"onUpdate:modelValue":r[5]||(r[5]=u=>g.value.total_time=u),precision:0,min:1,placeholder:s.$t("chat.form.total_time.placeholder"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1}),a(E,{span:8},{default:e(()=>[a(C,{field:"status",label:s.$t("chat.form.status")},{default:e(()=>[a(te,{modelValue:g.value.status,"onUpdate:modelValue":r[6]||(r[6]=u=>g.value.status=u),options:m(fe),placeholder:s.$t("chat.form.selectDefault"),"allow-clear":""},null,8,["modelValue","options","placeholder"])]),_:1},8,["label"])]),_:1}),a(E,{span:8},{default:e(()=>[a(C,{field:"req_time",label:s.$t("chat.form.req_time")},{default:e(()=>[a(le,{modelValue:g.value.req_time,"onUpdate:modelValue":r[7]||(r[7]=u=>g.value.req_time=u),placeholder:["\u5F00\u59CB\u65F6\u95F4","\u7ED3\u675F\u65F6\u95F4"],"time-picker-props":{defaultValue:["00:00:00","23:59:59"]},"show-time":""},null,8,["modelValue"])]),_:1},8,["label"])]),_:1})]),_:1})]),_:1},8,["model"])]),_:1}),a(De,{style:{height:"84px"},direction:"vertical"}),a(E,{flex:"86px",style:{"text-align":"right"}},{default:e(()=>[a(j,{direction:"vertical",size:18},{default:e(()=>[a(P,{type:"primary",onClick:X},{icon:e(()=>[a(Xe)]),default:e(()=>[i(" "+n(s.$t("chat.form.search")),1)]),_:1}),a(P,{onClick:Te},{icon:e(()=>[a(Ce)]),default:e(()=>[i(" "+n(s.$t("chat.form.reset")),1)]),_:1})]),_:1})]),_:1})]),_:1}),a(De,{style:{"margin-top":"0","margin-bottom":"16px"}}),a(oe,{style:{"margin-bottom":"16px","align-items":"center"}},{default:e(()=>[a(E,{span:4},{default:e(()=>[a(j,null,{default:e(()=>[a(P,{type:"primary",onClick:r[8]||(r[8]=u=>ge({}))},{default:e(()=>[i(" \u5BFC\u51FA ")]),_:1}),k((t(),o(P,{type:"primary",status:"danger",onClick:r[9]||(r[9]=u=>Fe({action:"delete"}))},{default:e(()=>[i(" \u5220\u9664 ")]),_:1})),[[y,["admin"]]])]),_:1})]),_:1}),a(E,{span:12},{default:e(()=>[i(" \u82B1\u8D39 = ( \u63D0\u95EE \xD7 \u63D0\u95EE\u500D\u7387 + \u56DE\u7B54 \xD7 \u56DE\u7B54\u500D\u7387 ) \xF7 500000 \xA0\xA0\u6216\xA0\xA0 \u56DE\u7B54 \xF7 500000 ")]),_:1}),a(E,{span:3},{default:e(()=>[i(" RPM: \xA0"+n(ke.value.toLocaleString()),1)]),_:1}),a(E,{span:3},{default:e(()=>[i(" TPM: \xA0"+n(he.value.toLocaleString()),1)]),_:1}),a(E,{span:2,style:{display:"flex","align-items":"center","justify-content":"end"}},{default:e(()=>[a(ne,{content:s.$t("searchTable.actions.refresh")},{default:e(()=>[$("div",{class:"action-icon",onClick:X},[a(Ce,{size:"18"})])]),_:1},8,["content"]),a(tt,{onSelect:Ue},{content:e(()=>[(t(!0),_(G,null,K(m(Le),u=>(t(),o(et,{key:u.value,value:u.value,class:ht({active:u.value===Q.value})},{default:e(()=>[$("span",null,n(u.name),1)]),_:2},1032,["value","class"]))),128))]),default:e(()=>[a(ne,{content:s.$t("searchTable.actions.density")},{default:e(()=>[$("div",ma,[a(Ze,{size:"18"})])]),_:1},8,["content"])]),_:1}),a(ne,{content:s.$t("searchTable.actions.columnSetting")},{default:e(()=>[a(ut,{trigger:"click",position:"bl",onPopupVisibleChange:Oe},{content:e(()=>[$("div",pa,[(t(!0),_(G,null,K(U.value,(u,B)=>(t(),_("div",{key:u.dataIndex,class:"setting"},[$("div",fa,[a(lt)]),$("div",null,[a(ot,{modelValue:u.checked,"onUpdate:modelValue":se=>u.checked=se,onChange:se=>Me(se,u,B)},null,8,["modelValue","onUpdate:modelValue","onChange"])]),$("div",va,n(u.title==="#"?"\u5E8F\u5217\u53F7":u.title),1)]))),128))])]),default:e(()=>[$("div",ca,[a(at,{size:"18"})])]),_:1})]),_:1},8,["content"])]),_:1})]),_:1}),a(nt,{ref_key:"tableRef",ref:W,"row-key":"id",loading:m(c),pagination:M,columns:A.value,data:ce.value,bordered:!1,size:Q.value,"row-selection":Se,onPageChange:Ne,onPageSizeChange:Pe,onSelectionChange:Qe},{user_id:e(({record:u})=>[i(n(u.is_smart_match?"-":u.user_id),1)]),prompt_tokens:e(({record:u})=>[i(n(u.prompt_tokens?u.prompt_tokens:u.status===1&&u.billing_method===2?0:"-"),1)]),completion_tokens:e(({record:u})=>[i(n(u.completion_tokens?u.completion_tokens:u.status===1&&u.billing_method===2?0:"-"),1)]),total_tokens:e(({record:u})=>[i(n(u.total_tokens?`$${m(aa)(u.total_tokens)}`:u.status===1&&u.billing_method===2?0:"-"),1)]),stream:e(({record:u})=>[i(n(s.$t(`chat.dict.stream.${u.stream||!1}`)),1)]),conn_time:e(({record:u})=>[u.conn_time>3e4?k((t(),o(d,{key:0,color:"red"},{default:e(()=>[i(n(u.conn_time),1)]),_:2},1024)),[[y,["user"]]]):u.conn_time>15e3?k((t(),o(d,{key:1,color:"orange"},{default:e(()=>[i(n(u.conn_time),1)]),_:2},1024)),[[y,["user"]]]):u.conn_time>5e3?k((t(),o(d,{key:2,color:"gold"},{default:e(()=>[i(n(u.conn_time),1)]),_:2},1024)),[[y,["user"]]]):k((t(),o(d,{key:3,color:"green"},{default:e(()=>[i(n(u.conn_time||"-"),1)]),_:2},1024)),[[y,["user"]]]),u.conn_time>1e4?k((t(),o(d,{key:4,color:"red"},{default:e(()=>[i(n(u.conn_time),1)]),_:2},1024)),[[y,["admin"]]]):u.conn_time>5e3?k((t(),o(d,{key:5,color:"orange"},{default:e(()=>[i(n(u.conn_time),1)]),_:2},1024)),[[y,["admin"]]]):u.conn_time>3e3?k((t(),o(d,{key:6,color:"gold"},{default:e(()=>[i(n(u.conn_time),1)]),_:2},1024)),[[y,["admin"]]]):k((t(),o(d,{key:7,color:"green"},{default:e(()=>[i(n(u.conn_time||"-"),1)]),_:2},1024)),[[y,["admin"]]])]),duration:e(({record:u})=>[u.duration>18e4?k((t(),o(d,{key:0,color:"red"},{default:e(()=>[i(n(u.duration),1)]),_:2},1024)),[[y,["user"]]]):u.duration>12e4?k((t(),o(d,{key:1,color:"orange"},{default:e(()=>[i(n(u.duration),1)]),_:2},1024)),[[y,["user"]]]):u.duration>9e4?k((t(),o(d,{key:2,color:"gold"},{default:e(()=>[i(n(u.duration),1)]),_:2},1024)),[[y,["user"]]]):k((t(),o(d,{key:3,color:"green"},{default:e(()=>[i(n(u.duration||"-"),1)]),_:2},1024)),[[y,["user"]]]),u.duration>12e4?k((t(),o(d,{key:4,color:"red"},{default:e(()=>[i(n(u.duration),1)]),_:2},1024)),[[y,["admin"]]]):u.duration>9e4?k((t(),o(d,{key:5,color:"orange"},{default:e(()=>[i(n(u.duration),1)]),_:2},1024)),[[y,["admin"]]]):u.duration>6e4?k((t(),o(d,{key:6,color:"gold"},{default:e(()=>[i(n(u.duration),1)]),_:2},1024)),[[y,["admin"]]]):k((t(),o(d,{key:7,color:"green"},{default:e(()=>[i(n(u.duration||"-"),1)]),_:2},1024)),[[y,["admin"]]])]),total_time:e(({record:u})=>[u.total_time>18e4?k((t(),o(d,{key:0,color:"red"},{default:e(()=>[i(n(u.total_time),1)]),_:2},1024)),[[y,["user"]]]):u.total_time>12e4?k((t(),o(d,{key:1,color:"orange"},{default:e(()=>[i(n(u.total_time),1)]),_:2},1024)),[[y,["user"]]]):u.total_time>9e4?k((t(),o(d,{key:2,color:"gold"},{default:e(()=>[i(n(u.total_time),1)]),_:2},1024)),[[y,["user"]]]):k((t(),o(d,{key:3,color:"green"},{default:e(()=>[i(n(u.total_time||"-"),1)]),_:2},1024)),[[y,["user"]]]),u.total_time>12e4?k((t(),o(d,{key:4,color:"red"},{default:e(()=>[i(n(u.total_time),1)]),_:2},1024)),[[y,["admin"]]]):u.total_time>9e4?k((t(),o(d,{key:5,color:"orange"},{default:e(()=>[i(n(u.total_time),1)]),_:2},1024)),[[y,["admin"]]]):u.total_time>6e4?k((t(),o(d,{key:6,color:"gold"},{default:e(()=>[i(n(u.total_time),1)]),_:2},1024)),[[y,["admin"]]]):k((t(),o(d,{key:7,color:"green"},{default:e(()=>[i(n(u.total_time||"-"),1)]),_:2},1024)),[[y,["admin"]]])]),internal_time:e(({record:u})=>[u.internal_time>1e3?k((t(),o(d,{key:0,color:"red"},{default:e(()=>[i(n(u.internal_time),1)]),_:2},1024)),[[y,["user"]]]):u.internal_time>500?k((t(),o(d,{key:1,color:"orange"},{default:e(()=>[i(n(u.internal_time),1)]),_:2},1024)),[[y,["user"]]]):u.internal_time>300?k((t(),o(d,{key:2,color:"gold"},{default:e(()=>[i(n(u.internal_time),1)]),_:2},1024)),[[y,["user"]]]):k((t(),o(d,{key:3,color:"green"},{default:e(()=>[i(n(u.internal_time||"-"),1)]),_:2},1024)),[[y,["user"]]]),u.internal_time>500?k((t(),o(d,{key:4,color:"red"},{default:e(()=>[i(n(u.internal_time),1)]),_:2},1024)),[[y,["admin"]]]):u.internal_time>300?k((t(),o(d,{key:5,color:"orange"},{default:e(()=>[i(n(u.internal_time),1)]),_:2},1024)),[[y,["admin"]]]):u.internal_time>100?k((t(),o(d,{key:6,color:"gold"},{default:e(()=>[i(n(u.internal_time),1)]),_:2},1024)),[[y,["admin"]]]):k((t(),o(d,{key:7,color:"green"},{default:e(()=>[i(n(u.internal_time||"-"),1)]),_:2},1024)),[[y,["admin"]]])]),status:e(({record:u})=>[u.status===-1?(t(),o(d,{key:0,color:"red"},{default:e(()=>[i(n(s.$t(`chat.dict.status.${u.status}`)),1)]),_:2},1024)):u.status===2?(t(),o(d,{key:1,color:"gold"},{default:e(()=>[i(n(s.$t(`chat.dict.status.${u.status}`)),1)]),_:2},1024)):u.status===3?(t(),o(d,{key:2,color:"orange"},{default:e(()=>[i(n(s.$t(`chat.dict.status.${u.status}`)),1)]),_:2},1024)):(t(),o(d,{key:3,color:"green"},{default:e(()=>[i(n(s.$t(`chat.dict.status.${u.status}`)),1)]),_:2},1024))]),operations:e(({record:u})=>[a(P,{type:"text",size:"small",onClick:B=>Ye(u.id)},{default:e(()=>[i(n(s.$t("chat.columns.operations.view")),1)]),_:2},1032,["onClick"])]),_:1},8,["loading","pagination","columns","data","size","row-selection"]),a(st,{title:"\u65E5\u5FD7\u8BE6\u60C5",visible:Z.value,width:700,footer:!1,"unmount-on-close":"","render-to-body":"",onCancel:He},{default:e(()=>[$("div",ya,[k((t(),o(J,{column:2,bordered:""},{default:e(()=>[a(v,{label:"Trace ID",span:2},{default:e(()=>[m(c)?(t(),o(f,{key:0,animation:!0},{default:e(()=>[a(p,{rows:1})]),_:1})):(t(),_("span",ka,[i(n(l.value.trace_id)+" ",1),a(T,{class:"copy-btn",onClick:r[10]||(r[10]=u=>z(l.value.trace_id))})]))]),_:1}),a(v,{label:"\u8C03\u7528\u5BC6\u94A5",span:2},{default:e(()=>[m(c)?(t(),o(f,{key:0,animation:!0},{default:e(()=>[a(p,{rows:1})]),_:1})):(t(),_("span",ha,[i(n(l.value.creator)+" ",1),a(T,{class:"copy-btn",onClick:r[11]||(r[11]=u=>z(l.value.creator))})]))]),_:1}),a(v,{label:"\u7528\u6237ID"},{default:e(()=>[m(c)?(t(),o(f,{key:0,animation:!0},{default:e(()=>[a(p,{widths:["200px"],rows:1})]),_:1})):(t(),_("span",ba,n(l.value.user_id||"-"),1))]),_:1}),a(v,{label:"\u5E94\u7528ID"},{default:e(()=>[m(c)?(t(),o(f,{key:0,animation:!0},{default:e(()=>[a(p,{widths:["200px"],rows:1})]),_:1})):(t(),_("span",ga,n(l.value.app_id||"-"),1))]),_:1}),a(v,{label:"\u516C\u53F8"},{default:e(()=>[m(c)?(t(),o(f,{key:0,animation:!0},{default:e(()=>[a(p,{rows:1})]),_:1})):(t(),_("span",wa,n(l.value.corp_name),1))]),_:1}),a(v,{label:"\u6A21\u578B"},{default:e(()=>[m(c)?(t(),o(f,{key:0,animation:!0},{default:e(()=>[a(p,{rows:1})]),_:1})):(t(),_("span",Fa,n(l.value.model||"-"),1))]),_:1}),a(v,{label:"\u6A21\u578B\u7C7B\u578B"},{default:e(()=>[m(c)?(t(),o(f,{key:0,animation:!0},{default:e(()=>[a(p,{rows:1})]),_:1})):(t(),_("span",Ea,n(s.$t(`chat.dict.type.${l.value.type}`)),1))]),_:1}),a(v,{label:"\u6D41\u5F0F"},{default:e(()=>[m(c)?(t(),o(f,{key:0,animation:!0},{default:e(()=>[a(p,{rows:1})]),_:1})):(t(),_("span",Da,n(s.$t(`chat.dict.stream.${l.value.stream||!1}`)),1))]),_:1}),a(v,{label:"\u63D0\u95EE",span:2},{default:e(()=>[m(c)?(t(),o(f,{key:0,animation:!0},{default:e(()=>[a(p,{rows:1})]),_:1})):(t(),_("span",Ca,n(l.value.prompt||"-"),1))]),_:1}),a(v,{label:"\u56DE\u7B54",span:2},{default:e(()=>[m(c)?(t(),o(f,{key:0,animation:!0},{default:e(()=>[a(p,{rows:1})]),_:1})):(t(),_("span",$a,n(l.value.completion||"-"),1))]),_:1}),a(v,{label:"\u8BA1\u8D39\u65B9\u5F0F"},{default:e(()=>[m(c)?(t(),o(f,{key:0,animation:!0},{default:e(()=>[a(p,{rows:1})]),_:1})):(t(),_("span",Ba,n(s.$t(`chat.dict.billing_method.${l.value.text_quota.billing_method}`)),1))]),_:1}),a(v,{label:"\u82B1\u8D39\u4EE4\u724C\u6570"},{default:e(()=>[m(c)?(t(),o(f,{key:0,animation:!0},{default:e(()=>[a(p,{rows:1})]),_:1})):(t(),_("span",Aa,n(l.value.total_tokens?l.value.total_tokens:l.value.status===1&&l.value.text_quota.billing_method===2?0:"-"),1))]),_:1}),a(v,{label:"\u63D0\u95EE\u500D\u7387"},{default:e(()=>[m(c)?(t(),o(f,{key:0,animation:!0},{default:e(()=>[a(p,{rows:1})]),_:1})):(t(),_("span",qa,n(l.value.type!==100?l.value.text_quota.prompt_ratio||"-":l.value.multimodal_quota.text_quota.prompt_ratio||"-"),1))]),_:1}),a(v,{label:"\u56DE\u7B54\u500D\u7387"},{default:e(()=>[m(c)?(t(),o(f,{key:0,animation:!0},{default:e(()=>[a(p,{rows:1})]),_:1})):(t(),_("span",Ia,n(l.value.type!==100?l.value.text_quota.completion_ratio||"-":l.value.multimodal_quota.text_quota.completion_ratio||"-"),1))]),_:1}),a(v,{label:"\u63D0\u95EE\u4EE4\u724C\u6570"},{default:e(()=>[m(c)?(t(),o(f,{key:0,animation:!0},{default:e(()=>[a(p,{rows:1})]),_:1})):(t(),_("span",Va,n(l.value.prompt_tokens||"-"),1))]),_:1}),a(v,{label:"\u56DE\u7B54\u4EE4\u724C\u6570"},{default:e(()=>[m(c)?(t(),o(f,{key:0,animation:!0},{default:e(()=>[a(p,{rows:1})]),_:1})):(t(),_("span",xa,n(l.value.completion_tokens||"-"),1))]),_:1}),a(v,{label:"\u8FDE\u63A5\u8017\u65F6"},{default:e(()=>[m(c)?(t(),o(f,{key:0,animation:!0},{default:e(()=>[a(p,{rows:1})]),_:1})):(t(),_("span",Sa,[l.value.conn_time>3e4?(t(),o(d,{key:0,color:"red"},{default:e(()=>[i(n(l.value.conn_time)+" ms ",1)]),_:1})):l.value.conn_time>15e3?(t(),o(d,{key:1,color:"orange"},{default:e(()=>[i(n(l.value.conn_time)+" ms ",1)]),_:1})):l.value.conn_time>5e3?(t(),o(d,{key:2,color:"gold"},{default:e(()=>[i(n(l.value.conn_time)+" ms ",1)]),_:1})):(t(),o(d,{key:3,color:"green"},{default:e(()=>[i(n(l.value.conn_time||"-")+" ms",1)]),_:1}))]))]),_:1}),a(v,{label:"\u6301\u7EED\u65F6\u957F"},{default:e(()=>[m(c)?(t(),o(f,{key:0,animation:!0},{default:e(()=>[a(p,{rows:1})]),_:1})):(t(),_("span",za,[l.value.duration>18e4?(t(),o(d,{key:0,color:"red"},{default:e(()=>[i(n(l.value.duration)+" ms ",1)]),_:1})):l.value.duration>12e4?(t(),o(d,{key:1,color:"orange"},{default:e(()=>[i(n(l.value.duration)+" ms ",1)]),_:1})):l.value.duration>9e4?(t(),o(d,{key:2,color:"gold"},{default:e(()=>[i(n(l.value.duration)+" ms ",1)]),_:1})):(t(),o(d,{key:3,color:"green"},{default:e(()=>[i(n(l.value.duration||"-")+" ms",1)]),_:1}))]))]),_:1}),a(v,{label:"\u603B\u8017\u65F6"},{default:e(()=>[m(c)?(t(),o(f,{key:0,animation:!0},{default:e(()=>[a(p,{rows:1})]),_:1})):(t(),_("span",La,[l.value.total_time>18e4?(t(),o(d,{key:0,color:"red"},{default:e(()=>[i(n(l.value.total_time)+" ms ",1)]),_:1})):l.value.total_time>12e4?(t(),o(d,{key:1,color:"orange"},{default:e(()=>[i(n(l.value.total_time)+" ms ",1)]),_:1})):l.value.total_time>9e4?(t(),o(d,{key:2,color:"gold"},{default:e(()=>[i(n(l.value.total_time)+" ms ",1)]),_:1})):(t(),o(d,{key:3,color:"green"},{default:e(()=>[i(n(l.value.total_time||"-")+" ms",1)]),_:1}))]))]),_:1}),a(v,{label:"\u7ED3\u679C"},{default:e(()=>[m(c)?(t(),o(f,{key:0,animation:!0},{default:e(()=>[a(p,{rows:1})]),_:1})):(t(),_("span",Na,[l.value.status===1?(t(),o(d,{key:0,color:"green"},{default:e(()=>[i(n(s.$t(`chat.dict.status.${l.value.status}`)),1)]),_:1})):l.value.status===2?(t(),o(d,{key:1,color:"gold"},{default:e(()=>[i(n(s.$t(`chat.dict.status.${l.value.status}`)),1)]),_:1})):l.value.status===3?(t(),o(d,{key:2,color:"orange"},{default:e(()=>[i(n(s.$t(`chat.dict.status.${l.value.status}`)),1)]),_:1})):(t(),o(d,{key:3,color:"red"},{default:e(()=>[i(n(s.$t(`chat.dict.status.${l.value.status}`)),1)]),_:1}))]))]),_:1}),a(v,{label:"\u5BA2\u6237\u7AEFIP"},{default:e(()=>[m(c)?(t(),o(f,{key:0,animation:!0},{default:e(()=>[a(p,{widths:["200px"],rows:1})]),_:1})):(t(),_("span",Pa,n(l.value.client_ip||"-"),1))]),_:1}),a(v,{label:"\u8BF7\u6C42\u65F6\u95F4"},{default:e(()=>[m(c)?(t(),o(f,{key:0,animation:!0},{default:e(()=>[a(p,{widths:["200px"],rows:1})]),_:1})):(t(),_("span",Ta,n(l.value.req_time||"-"),1))]),_:1}),a(v,{label:"\u9519\u8BEF\u4FE1\u606F",span:2},{default:e(()=>[m(c)?(t(),o(f,{key:0,animation:!0},{default:e(()=>[a(p,{rows:1})]),_:1})):(t(),_("span",Ua,n(l.value.err_msg||"-"),1))]),_:1})]),_:1})),[[y,["user"]]]),k((t(),o(J,{column:2,bordered:""},{default:e(()=>[a(v,{label:"Trace ID",span:2},{default:e(()=>[m(c)?(t(),o(f,{key:0,animation:!0},{default:e(()=>[a(p,{rows:1})]),_:1})):(t(),_("span",Ma,[i(n(l.value.trace_id)+" ",1),a(T,{class:"copy-btn",onClick:r[12]||(r[12]=u=>z(l.value.trace_id))})]))]),_:1}),a(v,{label:"\u8C03\u7528\u5BC6\u94A5",span:2},{default:e(()=>[m(c)?(t(),o(f,{key:0,animation:!0},{default:e(()=>[a(p,{rows:1})]),_:1})):(t(),_("span",Oa,[i(n(l.value.is_smart_match?"-":l.value.creator)+" ",1),l.value.is_smart_match?bt("",!0):(t(),o(T,{key:0,class:"copy-btn",onClick:r[13]||(r[13]=u=>z(l.value.creator))}))]))]),_:1}),a(v,{label:"\u7528\u6237ID"},{default:e(()=>[m(c)?(t(),o(f,{key:0,animation:!0},{default:e(()=>[a(p,{widths:["200px"],rows:1})]),_:1})):(t(),_("span",Ra,n(l.value.is_smart_match?"-":l.value.user_id||"-"),1))]),_:1}),a(v,{label:"\u5E94\u7528ID"},{default:e(()=>[m(c)?(t(),o(f,{key:0,animation:!0},{default:e(()=>[a(p,{widths:["200px"],rows:1})]),_:1})):(t(),_("span",Ya,n(l.value.is_smart_match?"-":l.value.app_id||"-"),1))]),_:1}),a(v,{label:"\u516C\u53F8"},{default:e(()=>[m(c)?(t(),o(f,{key:0,animation:!0},{default:e(()=>[a(p,{rows:1})]),_:1})):(t(),_("span",Ha,n(l.value.corp_name),1))]),_:1}),a(v,{label:"\u6A21\u578B\u7C7B\u578B"},{default:e(()=>[m(c)?(t(),o(f,{key:0,animation:!0},{default:e(()=>[a(p,{rows:1})]),_:1})):(t(),_("span",ja,n(s.$t(`chat.dict.type.${l.value.type}`)),1))]),_:1}),a(v,{label:"\u8BF7\u6C42\u6A21\u578B\u540D\u79F0"},{default:e(()=>[m(c)?(t(),o(f,{key:0,animation:!0},{default:e(()=>[a(p,{widths:["200px"],rows:1})]),_:1})):(t(),_("span",Ja,n(l.value.name||"-"),1))]),_:1}),a(v,{label:"\u8BF7\u6C42\u6A21\u578B"},{default:e(()=>[m(c)?(t(),o(f,{key:0,animation:!0},{default:e(()=>[a(p,{rows:1})]),_:1})):(t(),_("span",Ga,n(l.value.model||"-"),1))]),_:1}),a(v,{label:"\u771F\u5B9E\u6A21\u578B\u540D\u79F0"},{default:e(()=>[m(c)?(t(),o(f,{key:0,animation:!0},{default:e(()=>[a(p,{rows:1})]),_:1})):(t(),_("span",Ka,n(l.value.real_model_name),1))]),_:1}),a(v,{label:"\u771F\u5B9E\u6A21\u578B"},{default:e(()=>[m(c)?(t(),o(f,{key:0,animation:!0},{default:e(()=>[a(p,{rows:1})]),_:1})):(t(),_("span",Qa,n(l.value.real_model),1))]),_:1}),a(v,{label:"\u542F\u7528\u540E\u5907"},{default:e(()=>[m(c)?(t(),o(f,{key:0,animation:!0},{default:e(()=>[a(p,{rows:1})]),_:1})):(t(),_("span",Wa,n(s.$t(`chat.dict.is_enable_fallback.${l.value.is_enable_fallback||!1}`)),1))]),_:1}),a(v,{label:"\u540E\u5907\u6A21\u578B"},{default:e(()=>{var u,B;return[m(c)?(t(),o(f,{key:0,animation:!0},{default:e(()=>[a(p,{widths:["200px"],rows:1})]),_:1})):(t(),_("span",Xa,n(((B=(u=l.value)==null?void 0:u.fallback_config)==null?void 0:B.fallback_model)||"-"),1))]}),_:1}),a(v,{label:"\u542F\u7528\u8F6C\u53D1"},{default:e(()=>[m(c)?(t(),o(f,{key:0,animation:!0},{default:e(()=>[a(p,{rows:1})]),_:1})):(t(),_("span",Za,n(s.$t(`chat.dict.is_enable_forward.${l.value.is_enable_forward||!1}`)),1))]),_:1}),a(v,{label:"\u8F6C\u53D1\u89C4\u5219"},{default:e(()=>[m(c)?(t(),o(f,{key:0,animation:!0},{default:e(()=>[a(p,{rows:1})]),_:1})):(t(),_("span",el,n(l.value.is_enable_forward?s.$t(`chat.dict.forward_rule.${l.value.forward_config.forward_rule||"1"}`):"-"),1))]),_:1}),a(v,{label:"\u542F\u7528\u4EE3\u7406"},{default:e(()=>[m(c)?(t(),o(f,{key:0,animation:!0},{default:e(()=>[a(p,{rows:1})]),_:1})):(t(),_("span",tl,n(s.$t(`chat.dict.is_enable_model_agent.${l.value.is_enable_model_agent||!1}`)),1))]),_:1}),a(v,{label:"\u4EE3\u7406\u540D\u79F0"},{default:e(()=>{var u,B;return[m(c)?(t(),o(f,{key:0,animation:!0},{default:e(()=>[a(p,{widths:["200px"],rows:1})]),_:1})):(t(),_("span",al,n(((B=(u=l.value)==null?void 0:u.model_agent)==null?void 0:B.name)||"-"),1))]}),_:1}),a(v,{label:"\u5BC6\u94A5",span:2},{default:e(()=>[m(c)?(t(),o(f,{key:0,animation:!0},{default:e(()=>[a(p,{rows:1})]),_:1})):(t(),_("span",ll,[i(n(l.value.key?l.value.key.length>0?l.value.key.substr(0,l.value.key.length/2>10?10:l.value.key.length/2)+"************************************"+l.value.key.substr(-(l.value.key.length/2>5?5:l.value.key.length/2)):l.value.key:"-")+" ",1),a(T,{class:"copy-btn",onClick:r[14]||(r[14]=u=>z(l.value.key))})]))]),_:1}),a(v,{label:"\u63D0\u95EE",span:2},{default:e(()=>[m(c)?(t(),o(f,{key:0,animation:!0},{default:e(()=>[a(p,{rows:1})]),_:1})):(t(),_("span",ol,n(l.value.prompt||"-"),1))]),_:1}),a(v,{label:"\u56DE\u7B54",span:2},{default:e(()=>[m(c)?(t(),o(f,{key:0,animation:!0},{default:e(()=>[a(p,{rows:1})]),_:1})):(t(),_("span",ul,n(l.value.completion||"-"),1))]),_:1}),a(v,{label:"\u8BA1\u8D39\u65B9\u5F0F"},{default:e(()=>[m(c)?(t(),o(f,{key:0,animation:!0},{default:e(()=>[a(p,{rows:1})]),_:1})):(t(),_("span",nl,n(s.$t(`chat.dict.billing_method.${l.value.text_quota.billing_method}`)),1))]),_:1}),a(v,{label:"\u82B1\u8D39\u4EE4\u724C\u6570"},{default:e(()=>[m(c)?(t(),o(f,{key:0,animation:!0},{default:e(()=>[a(p,{rows:1})]),_:1})):(t(),_("span",sl,n(l.value.total_tokens?l.value.total_tokens:l.value.status===1&&l.value.text_quota.billing_method===2?0:"-"),1))]),_:1}),a(v,{label:"\u63D0\u95EE\u500D\u7387"},{default:e(()=>[m(c)?(t(),o(f,{key:0,animation:!0},{default:e(()=>[a(p,{rows:1})]),_:1})):(t(),_("span",il,n(l.value.type!==100?l.value.text_quota.prompt_ratio||"-":l.value.multimodal_quota.text_quota.prompt_ratio||"-"),1))]),_:1}),a(v,{label:"\u56DE\u7B54\u500D\u7387"},{default:e(()=>[m(c)?(t(),o(f,{key:0,animation:!0},{default:e(()=>[a(p,{rows:1})]),_:1})):(t(),_("span",dl,n(l.value.type!==100?l.value.text_quota.completion_ratio||"-":l.value.multimodal_quota.text_quota.completion_ratio||"-"),1))]),_:1}),a(v,{label:"\u63D0\u95EE\u4EE4\u724C\u6570"},{default:e(()=>[m(c)?(t(),o(f,{key:0,animation:!0},{default:e(()=>[a(p,{rows:1})]),_:1})):(t(),_("span",rl,n(l.value.prompt_tokens||"-"),1))]),_:1}),a(v,{label:"\u56DE\u7B54\u4EE4\u724C\u6570"},{default:e(()=>[m(c)?(t(),o(f,{key:0,animation:!0},{default:e(()=>[a(p,{rows:1})]),_:1})):(t(),_("span",_l,n(l.value.completion_tokens||"-"),1))]),_:1}),a(v,{label:"\u8FDE\u63A5\u8017\u65F6"},{default:e(()=>[m(c)?(t(),o(f,{key:0,animation:!0},{default:e(()=>[a(p,{rows:1})]),_:1})):(t(),_("span",ml,[l.value.conn_time>1e4?(t(),o(d,{key:0,color:"red"},{default:e(()=>[i(n(l.value.conn_time)+" ms ",1)]),_:1})):l.value.conn_time>5e3?(t(),o(d,{key:1,color:"orange"},{default:e(()=>[i(n(l.value.conn_time)+" ms ",1)]),_:1})):l.value.conn_time>3e3?(t(),o(d,{key:2,color:"gold"},{default:e(()=>[i(n(l.value.conn_time)+" ms ",1)]),_:1})):(t(),o(d,{key:3,color:"green"},{default:e(()=>[i(n(l.value.conn_time||"-")+" ms",1)]),_:1}))]))]),_:1}),a(v,{label:"\u6301\u7EED\u65F6\u957F"},{default:e(()=>[m(c)?(t(),o(f,{key:0,animation:!0},{default:e(()=>[a(p,{rows:1})]),_:1})):(t(),_("span",cl,[l.value.duration>12e4?(t(),o(d,{key:0,color:"red"},{default:e(()=>[i(n(l.value.duration)+" ms ",1)]),_:1})):l.value.duration>9e4?(t(),o(d,{key:1,color:"orange"},{default:e(()=>[i(n(l.value.duration)+" ms ",1)]),_:1})):l.value.duration>6e4?(t(),o(d,{key:2,color:"gold"},{default:e(()=>[i(n(l.value.duration)+" ms ",1)]),_:1})):(t(),o(d,{key:3,color:"green"},{default:e(()=>[i(n(l.value.duration||"-")+" ms",1)]),_:1}))]))]),_:1}),a(v,{label:"\u603B\u8017\u65F6"},{default:e(()=>[m(c)?(t(),o(f,{key:0,animation:!0},{default:e(()=>[a(p,{rows:1})]),_:1})):(t(),_("span",pl,[l.value.total_time>12e4?(t(),o(d,{key:0,color:"red"},{default:e(()=>[i(n(l.value.total_time)+" ms ",1)]),_:1})):l.value.total_time>9e4?(t(),o(d,{key:1,color:"orange"},{default:e(()=>[i(n(l.value.total_time)+" ms ",1)]),_:1})):l.value.total_time>6e4?(t(),o(d,{key:2,color:"gold"},{default:e(()=>[i(n(l.value.total_time)+" ms ",1)]),_:1})):(t(),o(d,{key:3,color:"green"},{default:e(()=>[i(n(l.value.total_time||"-")+" ms",1)]),_:1}))]))]),_:1}),a(v,{label:"\u5185\u8017"},{default:e(()=>[m(c)?(t(),o(f,{key:0,animation:!0},{default:e(()=>[a(p,{rows:1})]),_:1})):(t(),_("span",fl,[l.value.internal_time>500?(t(),o(d,{key:0,color:"red"},{default:e(()=>[i(n(l.value.internal_time)+" ms ",1)]),_:1})):l.value.internal_time>300?(t(),o(d,{key:1,color:"orange"},{default:e(()=>[i(n(l.value.internal_time)+" ms ",1)]),_:1})):l.value.internal_time>100?(t(),o(d,{key:2,color:"gold"},{default:e(()=>[i(n(l.value.internal_time)+" ms ",1)]),_:1})):(t(),o(d,{key:3,color:"green"},{default:e(()=>[i(n(l.value.internal_time||"-")+" ms",1)]),_:1}))]))]),_:1}),a(v,{label:"\u7ED3\u679C"},{default:e(()=>[m(c)?(t(),o(f,{key:0,animation:!0},{default:e(()=>[a(p,{rows:1})]),_:1})):(t(),_("span",vl,[l.value.status===1?(t(),o(d,{key:0,color:"green"},{default:e(()=>[i(n(s.$t(`chat.dict.status.${l.value.status}`)),1)]),_:1})):l.value.status===2?(t(),o(d,{key:1,color:"gold"},{default:e(()=>[i(n(s.$t(`chat.dict.status.${l.value.status}`)),1)]),_:1})):l.value.status===3?(t(),o(d,{key:2,color:"orange"},{default:e(()=>[i(n(s.$t(`chat.dict.status.${l.value.status}`)),1)]),_:1})):(t(),o(d,{key:3,color:"red"},{default:e(()=>[i(n(s.$t(`chat.dict.status.${l.value.status}`)),1)]),_:1}))]))]),_:1}),a(v,{label:"\u672C\u5730IP"},{default:e(()=>[m(c)?(t(),o(f,{key:0,animation:!0},{default:e(()=>[a(p,{widths:["200px"],rows:1})]),_:1})):(t(),_("span",yl,n(l.value.local_ip||"-"),1))]),_:1}),a(v,{label:"\u5BA2\u6237\u7AEFIP"},{default:e(()=>[m(c)?(t(),o(f,{key:0,animation:!0},{default:e(()=>[a(p,{widths:["200px"],rows:1})]),_:1})):(t(),_("span",kl,n(l.value.client_ip||"-"),1))]),_:1}),a(v,{label:"\u8FDC\u7A0BIP"},{default:e(()=>[m(c)?(t(),o(f,{key:0,animation:!0},{default:e(()=>[a(p,{widths:["200px"],rows:1})]),_:1})):(t(),_("span",hl,n(l.value.remote_ip||"-"),1))]),_:1}),a(v,{label:"\u8BF7\u6C42\u65F6\u95F4"},{default:e(()=>[m(c)?(t(),o(f,{key:0,animation:!0},{default:e(()=>[a(p,{widths:["200px"],rows:1})]),_:1})):(t(),_("span",bl,n(l.value.req_time||"-"),1))]),_:1}),a(v,{label:"\u521B\u5EFA\u65F6\u95F4"},{default:e(()=>[m(c)?(t(),o(f,{key:0,animation:!0},{default:e(()=>[a(p,{widths:["200px"],rows:1})]),_:1})):(t(),_("span",gl,n(l.value.created_at||"-"),1))]),_:1}),a(v,{label:"\u9519\u8BEF\u4FE1\u606F",span:2},{default:e(()=>[m(c)?(t(),o(f,{key:0,animation:!0},{default:e(()=>[a(p,{rows:1})]),_:1})):(t(),_("span",wl,n(l.value.err_msg||"-"),1))]),_:1})]),_:1})),[[y,["admin"]]]),a(J,{layout:"inline-vertical",column:2,style:{"margin-top":"10px",position:"relative"}},{default:e(()=>[a(v,{span:2},{default:e(()=>[a(Be,{type:"card"},{default:e(()=>[a($e,{key:"1",title:"\u63D0\u95EE\u4E0A\u4E0B\u6587"},{default:e(()=>[m(c)?(t(),o(f,{key:0,animation:!0},{default:e(()=>[a(p,{rows:3})]),_:1})):(t(),o(j,{key:1},{default:e(()=>[l.value.messages?(t(),o(m(xe),{key:0,path:"res",data:l.value.messages,"show-length":!0},null,8,["data"])):(t(),_("span",Fl,"-"))]),_:1}))]),_:1})]),_:1})]),_:1})]),_:1}),k((t(),o(J,{layout:"inline-vertical",column:2,style:{"margin-top":"10px",position:"relative"}},{default:e(()=>[a(v,{span:2},{default:e(()=>[a(Be,{type:"card"},{default:e(()=>[a($e,{key:"1",title:"\u6A21\u578B\u4EE3\u7406"},{default:e(()=>[m(c)?(t(),o(f,{key:0,animation:!0},{default:e(()=>[a(p,{rows:3})]),_:1})):(t(),o(j,{key:1},{default:e(()=>[l.value.model_agent?(t(),o(m(xe),{key:0,data:l.value.model_agent,"show-length":!0},null,8,["data"])):(t(),_("span",El,"-"))]),_:1}))]),_:1})]),_:1})]),_:1})]),_:1})),[[y,["admin"]]])])]),_:1},8,["visible"]),a(Ae,{visible:L.value,"onUpdate:visible":r[16]||(r[16]=u=>L.value=u),title:s.$t("chat.form.title.chat_export"),onCancel:Je,onBeforeOk:je},{default:e(()=>[a(ue,{ref_key:"chatExportForm",ref:be,model:Y.value},{default:e(()=>[a(C,{field:"req_time",label:s.$t("chat.form.req_time"),rules:[{required:!0,message:s.$t("chat.error.req_time.required")}]},{default:e(()=>[a(le,{modelValue:Y.value.req_time,"onUpdate:modelValue":r[15]||(r[15]=u=>Y.value.req_time=u),placeholder:["\u5F00\u59CB\u65F6\u95F4","\u7ED3\u675F\u65F6\u95F4"],"time-picker-props":{defaultValue:["00:00:00","23:59:59"]},"show-time":""},null,8,["modelValue"])]),_:1},8,["label","rules"])]),_:1},8,["model"])]),_:1},8,["visible","title"]),a(Ae,{visible:N.value,"onUpdate:visible":r[18]||(r[18]=u=>N.value=u),title:s.$t("chat.form.title.chat_del"),onCancel:Ke,onBeforeOk:Ge},{default:e(()=>[a(ue,{ref_key:"chatDelForm",ref:we,model:H.value},{default:e(()=>[a(C,{field:"value",label:s.$t("chat.form.req_time"),rules:[{required:!0,message:s.$t("chat.error.req_time.required")}]},{default:e(()=>[a(le,{modelValue:H.value.value,"onUpdate:modelValue":r[17]||(r[17]=u=>H.value.value=u),placeholder:["\u5F00\u59CB\u65F6\u95F4","\u7ED3\u675F\u65F6\u95F4"],"time-picker-props":{defaultValue:["00:00:00","23:59:59"]},"show-time":""},null,8,["modelValue"])]),_:1},8,["label","rules"])]),_:1},8,["model"])]),_:1},8,["visible","title"])]),_:1})])}}});const so=ft(Cl,[["__scopeId","data-v-10f59e92"]]);export{so as default};
