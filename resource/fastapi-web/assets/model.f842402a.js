import{u as Ie,C as Ve,o as ze,x as Se,I as De,y as Ae,_ as qe}from"./index.05df1f52.js";/* empty css               *//* empty css               *//* empty css              *//* empty css              *//* empty css                *//* empty css                *//* empty css               *//* empty css              *//* empty css                */import{c as F,S as Le}from"./sortable.esm.2109e0e3.js";/* empty css               *//* empty css               *//* empty css               *//* empty css               *//* empty css              */import{d as Be,r as X,e as _,c as S,w as Me,B as c,C as f,aH as t,aG as l,aL as u,aM as i,u as D,aJ as A,aI as q,aD as N,F as g,D as Te,n as Fe,aK as Ne,aF as Pe,bC as Ue,b2 as xe,bA as Ke,b1 as Ee,bB as Oe,bD as je,bE as Ge,b5 as Re,bF as He,ab as Je,aU as Xe,bi as Qe,a5 as We,bj as Ye,bl as Ze,bm as et,b4 as tt,bG as at,bH as lt,bI as ot,bJ as nt}from"./arco.aed15247.js";import{h as st}from"./vue.0cc5b64a.js";import{u as it}from"./loading.b5911e1d.js";import{s as ut,q as ct,a as dt}from"./key.f1df2528.js";import{q as rt}from"./model.89eea4c7.js";import{e as pt}from"./agent.4ef33138.js";import"./chart.9aa6eafa.js";import"./base.87fcf6e2.js";const mt={class:"container"},_t={class:"action-icon"},yt={class:"action-icon"},ft={id:"tableSetting"},gt={style:{"margin-right":"4px",cursor:"move"}},vt={class:"title"},kt={key:0,class:"circle red"},bt={key:1,class:"circle"},ht={name:"KeyList"},$t=Be({...ht,setup(wt){const P=st(),Q=X({type:"checkbox",showCheckedAll:!0,onlyCurrent:!1}),U=_([]);(async()=>{try{const{data:e}=await rt();U.value=e.items}catch{}})();const x=_([]);(async()=>{try{const{data:e}=await pt();x.value=e.items}catch{}})();const W=async e=>{b(!0);try{await ut(e),V()}catch{}finally{b(!1)}},K=()=>({type:2,corp:"",key:"",models:[],model_agents:[],quota:_(),status:_(),created_at:[]}),{loading:Y,setLoading:b}=it(!0),{t:n}=Ie(),E=_([]),s=_(K()),v=_([]),w=_([]),L=_("medium"),h={current:1,pageSize:10,showTotal:!0,showPageSize:!0},C=X({...h}),Z=S(()=>[{name:n("searchTable.size.mini"),value:"mini"},{name:n("searchTable.size.small"),value:"small"},{name:n("searchTable.size.medium"),value:"medium"},{name:n("searchTable.size.large"),value:"large"}]),ee=S(()=>[{title:n("key.columns.corp"),dataIndex:"corp",slotName:"corp",align:"center",width:110},{title:n("key.columns.key"),dataIndex:"key",slotName:"key",align:"center",ellipsis:!0,tooltip:!0},{title:n("key.columns.models"),dataIndex:"model_names",slotName:"model_names",align:"center",ellipsis:!0,tooltip:!0},{title:n("key.columns.quota"),dataIndex:"quota",slotName:"quota",align:"center",width:80},{title:n("key.columns.status"),dataIndex:"status",slotName:"status",align:"center",width:80},{title:n("key.columns.updated_at"),dataIndex:"updated_at",slotName:"updated_at",align:"center",width:170},{title:n("key.columns.operations"),dataIndex:"operations",slotName:"operations",align:"center",width:220}]),te=S(()=>[{label:n("key.dict.corp.OpenAI"),value:"OpenAI"},{label:n("key.dict.corp.Baidu"),value:"Baidu"},{label:n("key.dict.corp.Xfyun"),value:"Xfyun"},{label:n("key.dict.corp.Aliyun"),value:"Aliyun"},{label:n("key.dict.corp.GLM"),value:"GLM"},{label:n("key.dict.corp.Midjourney"),value:"Midjourney"}]),ae=S(()=>[{label:n("key.dict.status.1"),value:1},{label:n("key.dict.status.2"),value:2}]),O=new Array(0);P.query.agent_id&&(O[0]=P.query.agent_id);const I=async(e={...h,type:2,model_agents:O})=>{b(!0);try{const{data:o}=await ct(e);E.value=o.items,C.current=e.current,C.pageSize=e.pageSize,C.total=o.paging.total}catch{}finally{b(!1)}},V=()=>{I({...h,...s.value})},le=e=>{I({...h,...s.value,current:e})},oe=e=>{h.pageSize=e,I({...h,...s.value})};I();const ne=()=>{s.value=K()},se=async e=>{b(!0);try{e.status=e.status===1?2:1,await dt(e),V()}catch{}finally{b(!1)}},ie=(e,o)=>{L.value=e},ue=(e,o,p)=>{e?v.value.splice(p,0,o):v.value=w.value.filter(d=>d.dataIndex!==o.dataIndex)},j=(e,o,p,d=!1)=>{const m=d?F(e):e;return o>-1&&p>-1&&m.splice(o,1,m.splice(p,1,m[o]).pop()),m},ce=e=>{e&&Fe(()=>{const o=document.getElementById("tableSetting");new Le(o,{onEnd(p){const{oldIndex:d,newIndex:m}=p;j(v.value,d,m),j(w.value,d,m)}})})};return Me(()=>ee.value,e=>{v.value=F(e),v.value.forEach((o,p)=>{o.checked=!0}),w.value=F(v.value)},{deep:!0,immediate:!0}),(e,o)=>{const p=Ve,d=Ne,m=Pe,z=Ue,$=xe,r=Ke,de=Ee,G=Oe,re=je,B=Ge,pe=Re,R=He,me=Je,k=Xe,H=ze,J=Qe,_e=We,M=Ye,ye=Se,fe=Ze,ge=et,ve=De,ke=Ae,be=tt,he=at,$e=lt,we=ot,Ce=nt;return c(),f("div",mt,[t(m,{class:"container-breadcrumb"},{default:l(()=>[t(d,null,{default:l(()=>[t(p)]),_:1}),t(d,null,{default:l(()=>[u(i(e.$t("menu.key")),1)]),_:1}),t(d,null,{default:l(()=>[u(i(e.$t("menu.key.model.list")),1)]),_:1})]),_:1}),t(Ce,{class:"general-card",title:e.$t("menu.key.model.list"),bordered:!1,"header-style":{padding:"20px"},"body-style":{padding:"0 20px 20px"}},{default:l(()=>[t(B,null,{default:l(()=>[t(r,{flex:1},{default:l(()=>[t(pe,{model:s.value,"label-col-props":{span:5},"wrapper-col-props":{span:18},"label-align":"left"},{default:l(()=>[t(B,{gutter:16},{default:l(()=>[t(r,{span:8},{default:l(()=>[t($,{field:"corp",label:e.$t("key.form.corp")},{default:l(()=>[t(z,{modelValue:s.value.corp,"onUpdate:modelValue":o[0]||(o[0]=a=>s.value.corp=a),options:D(te),placeholder:e.$t("key.form.selectDefault"),"allow-clear":""},null,8,["modelValue","options","placeholder"])]),_:1},8,["label"])]),_:1}),t(r,{span:8},{default:l(()=>[t($,{field:"key",label:e.$t("key.form.key")},{default:l(()=>[t(de,{modelValue:s.value.key,"onUpdate:modelValue":o[1]||(o[1]=a=>s.value.key=a),placeholder:e.$t("key.form.key.placeholder"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1}),t(r,{span:8},{default:l(()=>[t($,{field:"models",label:e.$t("key.form.models")},{default:l(()=>[t(z,{modelValue:s.value.models,"onUpdate:modelValue":o[2]||(o[2]=a=>s.value.models=a),placeholder:e.$t("key.form.selectDefault"),"max-tag-count":3,multiple:"","allow-search":"","allow-clear":""},{default:l(()=>[(c(!0),f(A,null,q(U.value,a=>(c(),N(G,{key:a.id,value:a.id,label:a.name},null,8,["value","label"]))),128))]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1}),t(r,{span:8},{default:l(()=>[t($,{field:"model_agents",label:e.$t("key.form.modelAgents")},{default:l(()=>[t(z,{modelValue:s.value.model_agents,"onUpdate:modelValue":o[3]||(o[3]=a=>s.value.model_agents=a),placeholder:e.$t("key.form.selectDefault"),"max-tag-count":3,multiple:"","allow-search":"","allow-clear":""},{default:l(()=>[(c(!0),f(A,null,q(x.value,a=>(c(),N(G,{key:a.id,value:a.id,label:a.name},null,8,["value","label"]))),128))]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label"])]),_:1}),t(r,{span:8},{default:l(()=>[t($,{field:"status",label:e.$t("key.form.status")},{default:l(()=>[t(z,{modelValue:s.value.status,"onUpdate:modelValue":o[4]||(o[4]=a=>s.value.status=a),options:D(ae),placeholder:e.$t("key.form.selectDefault"),"allow-clear":""},null,8,["modelValue","options","placeholder"])]),_:1},8,["label"])]),_:1}),t(r,{span:8},{default:l(()=>[t($,{field:"created_at",label:e.$t("key.form.created_at")},{default:l(()=>[t(re,{modelValue:s.value.created_at,"onUpdate:modelValue":o[5]||(o[5]=a=>s.value.created_at=a),style:{width:"100%"}},null,8,["modelValue"])]),_:1},8,["label"])]),_:1})]),_:1})]),_:1},8,["model"])]),_:1}),t(R,{style:{height:"84px"},direction:"vertical"}),t(r,{flex:"86px",style:{"text-align":"right"}},{default:l(()=>[t(J,{direction:"vertical",size:18},{default:l(()=>[t(k,{type:"primary",onClick:V},{icon:l(()=>[t(me)]),default:l(()=>[u(" "+i(e.$t("key.form.search")),1)]),_:1}),t(k,{onClick:ne},{icon:l(()=>[t(H)]),default:l(()=>[u(" "+i(e.$t("key.form.reset")),1)]),_:1})]),_:1})]),_:1})]),_:1}),t(R,{style:{"margin-top":"0"}}),t(B,{style:{"margin-bottom":"16px"}},{default:l(()=>[t(r,{span:12},{default:l(()=>[t(J,null,{default:l(()=>[t(k,{type:"primary",onClick:o[6]||(o[6]=a=>e.$router.push({name:"KeyCreate"}))},{icon:l(()=>[t(_e)]),default:l(()=>[u(" "+i(e.$t("key.operation.create")),1)]),_:1})]),_:1})]),_:1}),t(r,{span:12,style:{display:"flex","align-items":"center","justify-content":"end"}},{default:l(()=>[t(M,{content:e.$t("searchTable.actions.refresh")},{default:l(()=>[g("div",{class:"action-icon",onClick:V},[t(H,{size:"18"})])]),_:1},8,["content"]),t(ge,{onSelect:ie},{content:l(()=>[(c(!0),f(A,null,q(D(Z),a=>(c(),N(fe,{key:a.value,value:a.value,class:Te({active:a.value===L.value})},{default:l(()=>[g("span",null,i(a.name),1)]),_:2},1032,["value","class"]))),128))]),default:l(()=>[t(M,{content:e.$t("searchTable.actions.density")},{default:l(()=>[g("div",_t,[t(ye,{size:"18"})])]),_:1},8,["content"])]),_:1}),t(M,{content:e.$t("searchTable.actions.columnSetting")},{default:l(()=>[t(he,{trigger:"click",position:"bl",onPopupVisibleChange:ce},{content:l(()=>[g("div",ft,[(c(!0),f(A,null,q(w.value,(a,y)=>(c(),f("div",{key:a.dataIndex,class:"setting"},[g("div",gt,[t(ke)]),g("div",null,[t(be,{modelValue:a.checked,"onUpdate:modelValue":T=>a.checked=T,onChange:T=>ue(T,a,y)},null,8,["modelValue","onUpdate:modelValue","onChange"])]),g("div",vt,i(a.title==="#"?"\u5E8F\u5217\u53F7":a.title),1)]))),128))])]),default:l(()=>[g("div",yt,[t(ve,{size:"18"})])]),_:1})]),_:1},8,["content"])]),_:1})]),_:1}),t(we,{"row-key":"id",loading:D(Y),pagination:C,columns:v.value,data:E.value,bordered:!1,size:L.value,"row-selection":Q,onPageChange:le,onPageSizeChange:oe},{type:l(({record:a})=>[u(i(e.$t(`key.dict.type.${a.type}`)),1)]),corp:l(({record:a})=>[u(i(e.$t(`key.dict.corp.${a.corp}`)),1)]),model_names:l(({record:a})=>{var y;return[u(i(((y=a==null?void 0:a.model_names)==null?void 0:y.join(","))||"-"),1)]}),dataFormat:l(({record:a})=>[u(i(e.$t(`key.dict.data_format.${a.data_format}`)),1)]),quota:l(({record:a})=>[u(i((a==null?void 0:a.quota)||"-"),1)]),status:l(({record:a})=>[a.status===2?(c(),f("span",kt)):(c(),f("span",bt)),u(" "+i(e.$t(`key.dict.status.${a.status}`)),1)]),operations:l(({record:a})=>[t(k,{type:"text",size:"small",onClick:y=>e.$router.push({name:"KeyDetail",query:{id:`${a.id}`}})},{default:l(()=>[u(i(e.$t("key.columns.operations.view")),1)]),_:2},1032,["onClick"]),t(k,{type:"text",size:"small",onClick:y=>e.$router.push({name:"KeyUpdate",query:{id:`${a.id}`}})},{default:l(()=>[u(i(e.$t("key.columns.operations.update")),1)]),_:2},1032,["onClick"]),t(k,{type:"text",size:"small",onClick:y=>se({id:`${a.id}`,status:Number(`${a.status}`)})},{default:l(()=>[u(i(e.$t(`key.columns.operations.status.${a.status}`)),1)]),_:2},1032,["onClick"]),t($e,{content:"\u4F60\u786E\u5B9A\u8981\u5220\u9664\u5417?",onOk:y=>W({id:`${a.id}`})},{default:l(()=>[t(k,{type:"text",size:"small"},{default:l(()=>[u(i(e.$t("key.columns.operations.delete")),1)]),_:1})]),_:2},1032,["onOk"])]),_:1},8,["loading","pagination","columns","data","size","row-selection"])]),_:1},8,["title"])])}}});const Qt=qe($t,[["__scopeId","data-v-f3803ff5"]]);export{Qt as default};
