import{u as ne,_ as ee,i as Te}from"./index.c1e68822.js";/* empty css                *//* empty css                *//* empty css               *//* empty css               */import{d as le,e as E,B as a,C as w,aH as l,aG as r,aJ as P,aI as M,aD as n,F as ie,aE as _,aL as B,aM as U,aT as de,bN as _e,c3 as ce,aU as pe,bJ as me,bK as ge,bI as re,bC as ue,bD as oe,b1 as fe,b2 as ye,aS as be,b5 as ve,a_ as Fe,g as Ee,bu as Y,bv as Z,a5 as qe,a4 as Se,b4 as He,bi as Le,aK as Re,aF as ze,bb as Ge,b6 as Ie}from"./arco.a9260898.js";/* empty css               *//* empty css                *//* empty css                */import{u as he}from"./loading.1f346a94.js";import{s as Ce,a as $e,b as Be,c as ke}from"./sys_config.8d0e2a0f.js";/* empty css                *//* empty css               */import"./chart.d103b168.js";import"./vue.ad52ddbe.js";const Pe={class:"list-wrap",style:{"margin-top":"10px"}},Me={class:"card-wrap"},xe={name:"Global"},Je=le({...xe,setup(ae){const{proxy:b}=Ee(),{setLoading:f}=he(!0),{t:y}=ne(),v=E(!1),q=E(""),k=E(),e=E({email:{},http:{},core:{},debug:{}}),H=async u=>{q.value=y(`sys.config.item.title.${u.action}`),e.value.action=u.action,v.value=!0},L=async u=>{b.$modal.warning({title:"\u8B66\u544A",titleAlign:"center",content:`\u662F\u5426\u786E\u5B9A\u91CD\u7F6E${y(`sys.config.item.title.${u.action}`)}?`,hideCancel:!1,onOk:()=>{J(u)}})},G=async u=>{var s;if(await((s=k.value)==null?void 0:s.validate())){v.value=!0,u(!1);return}f(!0);try{await Ce(e.value),u(),b.$message.success("\u64CD\u4F5C\u6210\u529F"),h()}catch{u(!1)}finally{f(!1)}},x=()=>{v.value=!1},J=async u=>{f(!0);try{await $e({action:u.action}),b.$message.success("\u64CD\u4F5C\u6210\u529F"),h()}finally{f(!1)}},K=async u=>{f(!0);try{await Be({action:u.action,open:u.open||!1}),b.$message.success("\u64CD\u4F5C\u6210\u529F"),h()}finally{f(!1)}},N=E({}),h=async()=>{const{data:u}=await ke();e.value.email=u.email,e.value.http=u.http,e.value.core=u.core,e.value.debug=u.debug,N.value=[{action:"email",title:y("sys.config.item.title.email"),description:"\u914D\u7F6E\u53D1\u4FE1\u90AE\u7BB1\u4FE1\u606F, \u7528\u4E8E\u6CE8\u518C\u3001\u767B\u5F55\u3001\u53D1\u9001\u6D88\u606F\u7B49\u573A\u666F, \u5F53\u6709\u7AD9\u70B9\u914D\u7F6E\u4E14\u540C\u65F6\u914D\u7F6E\u53D1\u4FE1\u90AE\u7BB1\u4FE1\u606F\u65F6, \u7AD9\u70B9\u914D\u7F6E\u7684\u4F18\u5148",open:e.value.email.open,config:!0,reset:!0},{action:"http",title:y("sys.config.item.title.http"),description:"\u914D\u7F6EHTTP\u8BF7\u6C42\u8D85\u65F6\u65F6\u95F4\u548C\u4EE3\u7406\u5730\u5740",config:!0,reset:!0},{action:"core",title:y("sys.config.item.title.core"),description:"\u7CFB\u7EDF\u9996\u6B21\u4F7F\u7528\u65F6\u53EF\u6839\u636E\u5B9E\u9645\u60C5\u51B5\u4FEE\u6539, \u540E\u671F\u5982\u82E5\u518D\u6B21\u4FEE\u6539\u5BC6\u94A5\u524D\u7F00, \u5386\u53F2\u5E94\u7528\u5BC6\u94A5\u5C06\u65E0\u6CD5\u901A\u8FC7\u6838\u9A8C, \u8BF7\u8C28\u614E\u4FEE\u6539",config:!0,reset:!0},{action:"debug",title:y("sys.config.item.title.debug"),description:"\u8C03\u8BD5\u5F00\u5173\u6253\u5F00\u540E, \u65E5\u5FD7\u4F1A\u6253\u5370\u66F4\u591A\u8BE6\u7EC6\u4FE1\u606F, \u65E5\u5FD7\u7EA7\u522B(logger.level)\u9700\u914D\u7F6E\u4E3A: debug",open:e.value.debug.open}]};return h(),(u,c)=>{const s=de,g=_e,$=ce,O=pe,R=me,T=ge,o=re,m=ue,z=oe,A=fe,D=ye,F=be,I=ve,j=Fe;return a(),w("div",Pe,[l(z,{class:"list-row",gutter:24},{default:r(()=>[(a(!0),w(P,null,M(N.value,t=>(a(),n(m,{key:t.action,xs:12,sm:12,md:12,lg:6,xl:6,xxl:6,class:"list-col"},{default:r(()=>[ie("div",Me,[l(o,{title:t.title,bordered:!1,"header-style":{padding:"16px"},"body-style":{padding:"0px 16px"}},{extra:r(()=>[t.open!==void 0?(a(),n(s,{key:0,modelValue:t.open,"onUpdate:modelValue":S=>t.open=S,onChange:S=>K(t)},null,8,["modelValue","onUpdate:modelValue","onChange"])):_("",!0)]),actions:r(()=>[t.reset?(a(),n(O,{key:0,onClick:S=>L(t)},{default:r(()=>[B(U(u.$t("button.reset")),1)]),_:2},1032,["onClick"])):_("",!0),t.config?(a(),n(O,{key:1,type:"primary",onClick:S=>H(t)},{default:r(()=>[B(U(u.$t("button.config")),1)]),_:2},1032,["onClick"])):_("",!0)]),skeleton:r(()=>[l(T,{animation:!0},{default:r(()=>[l(R,{widths:["50%","50%","100%","40%"],rows:4}),l(R,{widths:["40%"],rows:1})]),_:1})]),default:r(()=>[l($,null,{description:r(()=>[B(U(t.description)+" ",1),l(g,{style:{"margin-top":"10px"},data:t.data,layout:"inline-horizontal",column:2},null,8,["data"])]),_:2},1024)]),_:2},1032,["title"])])]),_:2},1024))),128))]),_:1}),l(j,{visible:v.value,"onUpdate:visible":c[9]||(c[9]=t=>v.value=t),title:u.$t(q.value),onCancel:x,onBeforeOk:G},{default:r(()=>[l(I,{ref_key:"configForm",ref:k,model:e.value,"auto-label-width":""},{default:r(()=>[e.value.action==="email"?(a(),n(D,{key:0,field:"email.host",label:u.$t("sys.config.label.email.host"),rules:[{required:!0,message:u.$t("sys.config.error.email.host.required")}]},{default:r(()=>[l(A,{modelValue:e.value.email.host,"onUpdate:modelValue":c[0]||(c[0]=t=>e.value.email.host=t),placeholder:u.$t("sys.config.placeholder.email.host"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"])):_("",!0),e.value.action==="email"?(a(),n(D,{key:1,field:"email.port",label:u.$t("sys.config.label.email.port"),rules:[{required:!0,message:u.$t("sys.config.error.email.port.required")}]},{default:r(()=>[l(F,{modelValue:e.value.email.port,"onUpdate:modelValue":c[1]||(c[1]=t=>e.value.email.port=t),placeholder:u.$t("sys.config.placeholder.email.port"),min:1,"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"])):_("",!0),e.value.action==="email"?(a(),n(D,{key:2,field:"email.user_name",label:u.$t("sys.config.label.email.user_name"),rules:[{required:!0,message:u.$t("sys.config.error.email.user_name.required")}]},{default:r(()=>[l(A,{modelValue:e.value.email.user_name,"onUpdate:modelValue":c[2]||(c[2]=t=>e.value.email.user_name=t),placeholder:u.$t("sys.config.placeholder.email.user_name"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"])):_("",!0),e.value.action==="email"?(a(),n(D,{key:3,field:"email.password",label:u.$t("sys.config.label.email.password"),rules:[{required:!0,message:u.$t("sys.config.error.email.password.required")}]},{default:r(()=>[l(A,{modelValue:e.value.email.password,"onUpdate:modelValue":c[3]||(c[3]=t=>e.value.email.password=t),placeholder:u.$t("sys.config.placeholder.email.password"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"])):_("",!0),e.value.action==="email"?(a(),n(D,{key:4,field:"email.from_name",label:u.$t("sys.config.label.email.from_name"),rules:[{required:!0,message:u.$t("sys.config.error.email.from_name.required")}]},{default:r(()=>[l(A,{modelValue:e.value.email.from_name,"onUpdate:modelValue":c[4]||(c[4]=t=>e.value.email.from_name=t),placeholder:u.$t("sys.config.placeholder.email.from_name"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"])):_("",!0),e.value.action==="http"?(a(),n(D,{key:5,field:"http.timeout",label:u.$t("sys.config.label.http.timeout"),rules:[{required:!0,message:u.$t("sys.config.error.http.timeout.required")}]},{default:r(()=>[l(F,{modelValue:e.value.http.timeout,"onUpdate:modelValue":c[5]||(c[5]=t=>e.value.http.timeout=t),placeholder:u.$t("sys.config.placeholder.http.timeout"),min:1,"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"])):_("",!0),e.value.action==="http"?(a(),n(D,{key:6,field:"http.proxy_url",label:u.$t("sys.config.label.http.proxy_url")},{default:r(()=>[l(A,{modelValue:e.value.http.proxy_url,"onUpdate:modelValue":c[6]||(c[6]=t=>e.value.http.proxy_url=t),placeholder:u.$t("sys.config.placeholder.http.proxy_url"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label"])):_("",!0),e.value.action==="core"?(a(),n(D,{key:7,field:"core.secret_key_prefix",label:u.$t("sys.config.label.core.secret_key_prefix")},{default:r(()=>[l(A,{modelValue:e.value.core.secret_key_prefix,"onUpdate:modelValue":c[7]||(c[7]=t=>e.value.core.secret_key_prefix=t),placeholder:u.$t("sys.config.placeholder.core.secret_key_prefix"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label"])):_("",!0),e.value.action==="core"?(a(),n(D,{key:8,field:"core.error_prefix",label:u.$t("sys.config.label.core.error_prefix")},{default:r(()=>[l(A,{modelValue:e.value.core.error_prefix,"onUpdate:modelValue":c[8]||(c[8]=t=>e.value.core.error_prefix=t),placeholder:u.$t("sys.config.placeholder.core.error_prefix"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label"])):_("",!0)]),_:1},8,["model"])]),_:1},8,["visible","title"])])}}});const Ve=ee(Je,[["__scopeId","data-v-4ee626f1"]]),Ke={class:"list-wrap",style:{"margin-top":"10px"}},je={class:"card-wrap"},Qe={name:"General"},We=le({...Qe,setup(ae){const{proxy:b}=Ee(),{setLoading:f}=he(!0),{t:y}=ne(),v=E(!1),q=E(""),k=E(),e=E({user_login_register:{},user_shield_error:{},admin_login:{}}),H=async s=>{s.action==="user_shield_error"&&e.value.user_shield_error.errors.length===0&&u(),q.value=y(`sys.config.item.title.${s.action}`),e.value.action=s.action,v.value=!0},L=async s=>{b.$modal.warning({title:"\u8B66\u544A",titleAlign:"center",content:`\u662F\u5426\u786E\u5B9A\u91CD\u7F6E${y(`sys.config.item.title.${s.action}`)}?`,hideCancel:!1,onOk:()=>{J(s)}})},G=async s=>{var $;if(await(($=k.value)==null?void 0:$.validate())){v.value=!0,s(!1);return}f(!0);try{await Ce(e.value),s(),b.$message.success("\u64CD\u4F5C\u6210\u529F"),h()}catch{s(!1)}finally{f(!1)}},x=()=>{v.value=!1,e.value.user_shield_error.errors.length>0&&!e.value.user_shield_error.errors[e.value.user_shield_error.errors.length-1]&&c(e.value.user_shield_error.errors.length-1)},J=async s=>{f(!0);try{await $e({action:s.action}),b.$message.success("\u64CD\u4F5C\u6210\u529F"),h()}finally{f(!1)}},K=async s=>{f(!0);try{await Be({action:s.action,open:s.open||!1}),b.$message.success("\u64CD\u4F5C\u6210\u529F"),h()}finally{f(!1)}},N=E({}),h=async()=>{const{data:s}=await ke();e.value.user_login_register=s.user_login_register,e.value.user_shield_error=s.user_shield_error,e.value.admin_login=s.admin_login,N.value=[{action:"user_login_register",title:y("sys.config.item.title.user_login_register"),description:"\u914D\u7F6E\u767B\u5F55\u9875\u4E0A\u7684\u767B\u5F55\u65B9\u5F0F\u3001\u7528\u6237\u6CE8\u518C\u3001\u627E\u56DE\u5BC6\u7801\u4EE5\u53CA\u4F1A\u8BDD\u8FC7\u671F\u65F6\u957F, \u5BF9\u5E94\u7684\u5F00\u5173\u53EF\u63A7\u5236\u767B\u5F55\u9875\u4E0A\u5BF9\u5E94\u529F\u80FD\u7684\u663E\u793A, \u5173\u95ED\u7528\u6237\u6CE8\u518C\u65F6, \u901A\u8FC7\u90AE\u7BB1\u767B\u5F55\u4E5F\u65E0\u6CD5\u81EA\u52A8\u6CE8\u518C",config:!0,reset:!0},{action:"user_shield_error",title:y("sys.config.item.title.user_shield_error"),description:"\u7528\u6237\u67E5\u770B\u8C03\u7528\u65E5\u5FD7\u9519\u8BEF\u65F6, \u5305\u542B\u6709\u914D\u7F6E\u9519\u8BEF\u5185\u5BB9\u65F6\u5219\u5C4F\u853D\u663E\u793A, \u4E3A\u7A7A\u5219\u5C4F\u853D\u6240\u6709\u9519\u8BEF\u663E\u793A",open:e.value.user_shield_error.open,config:!0,reset:!0},{action:"admin_login",title:y("sys.config.item.title.admin_login"),description:"\u914D\u7F6E\u767B\u5F55\u9875\u4E0A\u7684\u767B\u5F55\u65B9\u5F0F\u3001\u627E\u56DE\u5BC6\u7801\u4EE5\u53CA\u4F1A\u8BDD\u8FC7\u671F\u65F6\u957F, \u5BF9\u5E94\u7684\u5F00\u5173\u53EF\u63A7\u5236\u767B\u5F55\u9875\u4E0A\u5BF9\u5E94\u529F\u80FD\u7684\u663E\u793A",config:!0,reset:!0}]};h();const u=()=>{e.value.user_shield_error.errors.push("")},c=s=>{e.value.user_shield_error.errors.splice(s,1)};return(s,g)=>{const $=de,O=_e,R=ce,T=pe,o=me,m=ge,z=re,A=ue,D=oe,F=ye,I=be,j=fe,t=qe,S=Se,De=ve,Q=Fe;return a(),w("div",Ke,[l(D,{class:"list-row",gutter:24},{default:r(()=>[(a(!0),w(P,null,M(N.value,i=>(a(),n(A,{key:i.action,xs:12,sm:12,md:12,lg:6,xl:6,xxl:6,class:"list-col"},{default:r(()=>[ie("div",je,[l(z,{title:i.title,bordered:!1,"header-style":{padding:"16px"},"body-style":{padding:"0px 16px"}},{extra:r(()=>[i.open!==void 0?(a(),n($,{key:0,modelValue:i.open,"onUpdate:modelValue":V=>i.open=V,onChange:V=>K(i)},null,8,["modelValue","onUpdate:modelValue","onChange"])):_("",!0)]),actions:r(()=>[i.reset?(a(),n(T,{key:0,onClick:V=>L(i)},{default:r(()=>[B(U(s.$t("button.reset")),1)]),_:2},1032,["onClick"])):_("",!0),i.config?(a(),n(T,{key:1,type:"primary",onClick:V=>H(i)},{default:r(()=>[B(U(s.$t("button.config")),1)]),_:2},1032,["onClick"])):_("",!0)]),skeleton:r(()=>[l(m,{animation:!0},{default:r(()=>[l(o,{widths:["50%","50%","100%","40%"],rows:4}),l(o,{widths:["40%"],rows:1})]),_:1})]),default:r(()=>[l(R,null,{description:r(()=>[B(U(i.description)+" ",1),l(O,{style:{"margin-top":"10px"},data:i.data,layout:"inline-horizontal",column:2},null,8,["data"])]),_:2},1024)]),_:2},1032,["title"])])]),_:2},1024))),128))]),_:1}),l(Q,{visible:v.value,"onUpdate:visible":g[10]||(g[10]=i=>v.value=i),title:s.$t(q.value),onCancel:x,onBeforeOk:G},{default:r(()=>[l(De,{ref_key:"configForm",ref:k,model:e.value,"auto-label-width":"",style:{"max-height":"300px"}},{default:r(()=>[e.value.action==="user_login_register"?(a(),n(F,{key:0,field:"user_login_register.account_login",label:s.$t("sys.config.label.user_login_register.account_login"),rules:[{required:!0}]},{default:r(()=>[l($,{modelValue:e.value.user_login_register.account_login,"onUpdate:modelValue":g[0]||(g[0]=i=>e.value.user_login_register.account_login=i)},null,8,["modelValue"])]),_:1},8,["label"])):_("",!0),e.value.action==="user_login_register"?(a(),n(F,{key:1,field:"user_login_register.email_login",label:s.$t("sys.config.label.user_login_register.email_login"),rules:[{required:!0}]},{default:r(()=>[l($,{modelValue:e.value.user_login_register.email_login,"onUpdate:modelValue":g[1]||(g[1]=i=>e.value.user_login_register.email_login=i)},null,8,["modelValue"])]),_:1},8,["label"])):_("",!0),e.value.action==="user_login_register"?(a(),n(F,{key:2,field:"user_login_register.email_register",label:s.$t("sys.config.label.user_login_register.email_register"),rules:[{required:!0}]},{default:r(()=>[l($,{modelValue:e.value.user_login_register.email_register,"onUpdate:modelValue":g[2]||(g[2]=i=>e.value.user_login_register.email_register=i)},null,8,["modelValue"])]),_:1},8,["label"])):_("",!0),e.value.action==="user_login_register"?(a(),n(F,{key:3,field:"user_login_register.email_retrieve",label:s.$t("sys.config.label.user_login_register.email_retrieve"),rules:[{required:!0}]},{default:r(()=>[l($,{modelValue:e.value.user_login_register.email_retrieve,"onUpdate:modelValue":g[3]||(g[3]=i=>e.value.user_login_register.email_retrieve=i)},null,8,["modelValue"])]),_:1},8,["label"])):_("",!0),e.value.action==="user_login_register"?(a(),n(F,{key:4,field:"user_login_register.session_expire",label:s.$t("sys.config.label.user_login_register.session_expire"),rules:[{required:!0,message:s.$t("sys.config.error.user_login_register.session_expire.required")}]},{default:r(()=>[l(I,{modelValue:e.value.user_login_register.session_expire,"onUpdate:modelValue":g[4]||(g[4]=i=>e.value.user_login_register.session_expire=i),placeholder:s.$t("sys.config.placeholder.user_login_register.session_expire"),min:10,"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"])):_("",!0),(a(!0),w(P,null,M(e.value.user_shield_error.errors,(i,V)=>Y((a(),n(F,{key:V,field:`user_shield_error.errors[${V}]`,label:`${V+1}. `+s.$t("sys.config.label.user_shield_error.errors"),rules:[{required:!0,message:s.$t("sys.config.error.user_shield_error.errors.required")}],"label-col-style":{padding:"0 16px 2px 0"}},{default:r(()=>[l(j,{modelValue:e.value.user_shield_error.errors[V],"onUpdate:modelValue":W=>e.value.user_shield_error.errors[V]=W,placeholder:s.$t("sys.config.placeholder.user_shield_error.errors"),"allow-clear":"",style:{width:"75%","margin-right":"5px"}},null,8,["modelValue","onUpdate:modelValue","placeholder"]),l(T,{type:"primary",shape:"circle",style:{margin:"0 10px 0 10px"},onClick:g[5]||(g[5]=W=>u())},{default:r(()=>[l(t)]),_:1}),l(T,{type:"secondary",shape:"circle",onClick:W=>c(V)},{default:r(()=>[l(S)]),_:2},1032,["onClick"])]),_:2},1032,["field","label","rules"])),[[Z,e.value.action==="user_shield_error"]])),128)),e.value.action==="admin_login"?(a(),n(F,{key:5,field:"admin_login.account_login",label:s.$t("sys.config.label.admin_login.account_login"),rules:[{required:!0}]},{default:r(()=>[l($,{modelValue:e.value.admin_login.account_login,"onUpdate:modelValue":g[6]||(g[6]=i=>e.value.admin_login.account_login=i)},null,8,["modelValue"])]),_:1},8,["label"])):_("",!0),e.value.action==="admin_login"?(a(),n(F,{key:6,field:"admin_login.email_login",label:s.$t("sys.config.label.admin_login.email_login"),rules:[{required:!0}]},{default:r(()=>[l($,{modelValue:e.value.admin_login.email_login,"onUpdate:modelValue":g[7]||(g[7]=i=>e.value.admin_login.email_login=i)},null,8,["modelValue"])]),_:1},8,["label"])):_("",!0),e.value.action==="admin_login"?(a(),n(F,{key:7,field:"admin_login.email_retrieve",label:s.$t("sys.config.label.admin_login.email_retrieve"),rules:[{required:!0}]},{default:r(()=>[l($,{modelValue:e.value.admin_login.email_retrieve,"onUpdate:modelValue":g[8]||(g[8]=i=>e.value.admin_login.email_retrieve=i)},null,8,["modelValue"])]),_:1},8,["label"])):_("",!0),e.value.action==="admin_login"?(a(),n(F,{key:8,field:"admin_login.session_expire",label:s.$t("sys.config.label.admin_login.session_expire"),rules:[{required:!0,message:s.$t("sys.config.error.admin_login.session_expire.required")}]},{default:r(()=>[l(I,{modelValue:e.value.admin_login.session_expire,"onUpdate:modelValue":g[9]||(g[9]=i=>e.value.admin_login.session_expire=i),placeholder:s.$t("sys.config.placeholder.admin_login.session_expire"),min:10,"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"])):_("",!0)]),_:1},8,["model"])]),_:1},8,["visible","title"])])}}});const Ae=ee(We,[["__scopeId","data-v-f1f30271"]]),Xe={class:"list-wrap",style:{"margin-top":"10px"}},Ye={class:"card-wrap"},Ze={name:"Task"},el=le({...Ze,setup(ae){const{proxy:b}=Ee(),{setLoading:f}=he(!0),{t:y}=ne(),v=E(!1),q=E(""),k=E(),e=E({statistics:{}}),H=async u=>{q.value=y(`sys.config.item.title.${u.action}`),e.value.action=u.action,v.value=!0},L=async u=>{b.$modal.warning({title:"\u8B66\u544A",titleAlign:"center",content:`\u662F\u5426\u786E\u5B9A\u91CD\u7F6E${y(`sys.config.item.title.${u.action}`)}?`,hideCancel:!1,onOk:()=>{J(u)}})},G=async u=>{var s;if(await((s=k.value)==null?void 0:s.validate())){v.value=!0,u(!1);return}f(!0);try{await Ce(e.value),u(),b.$message.success("\u64CD\u4F5C\u6210\u529F"),h()}catch{u(!1)}finally{f(!1)}},x=()=>{v.value=!1},J=async u=>{f(!0);try{await $e({action:u.action}),b.$message.success("\u64CD\u4F5C\u6210\u529F"),h()}finally{f(!1)}},K=async u=>{f(!0);try{await Be({action:u.action,open:u.open||!1}),b.$message.success("\u64CD\u4F5C\u6210\u529F"),h()}finally{f(!1)}},N=E({}),h=async()=>{const{data:u}=await ke();e.value.statistics=u.statistics,N.value=[{action:"statistics",title:y("sys.config.item.title.statistics"),description:"\u4EEA\u8868\u76D8\u4E0A\u5404\u7C7B\u6570\u636E\u4EE5\u53CA\u8D26\u5355\u660E\u7EC6\u7684\u7EDF\u8BA1\u4EFB\u52A1, \u7EDF\u8BA1\u95F4\u9694\u65F6\u95F4\u5EFA\u8BAE\u63A7\u5236\u572830\u5206\u949F\u4EE5\u5185, \u5355\u6B21\u5FAA\u73AF\u7EDF\u8BA1\u67E5\u8BE2\u6761\u6570\u5EFA\u8BAE\u63A7\u5236\u57281\u4E07\u4EE5\u5185, \u5355\u6B21\u7EDF\u8BA1\u4EFB\u52A1\u8D85\u65F6\u65F6\u95F4\u53EF\u6839\u636E\u5B9E\u9645\u60C5\u51B5\u914D\u7F6E, \u5EFA\u8BAE\u4E0D\u8981\u4F4E\u4E8E10\u5206\u949F",open:e.value.statistics.open,config:!0,reset:!0}]};return h(),(u,c)=>{const s=de,g=_e,$=ce,O=pe,R=me,T=ge,o=re,m=ue,z=oe,A=fe,D=ye,F=be,I=ve,j=Fe;return a(),w("div",Xe,[l(z,{class:"list-row",gutter:24},{default:r(()=>[(a(!0),w(P,null,M(N.value,t=>(a(),n(m,{key:t.action,xs:12,sm:12,md:12,lg:6,xl:6,xxl:6,class:"list-col"},{default:r(()=>[ie("div",Ye,[l(o,{title:t.title,bordered:!1,"header-style":{padding:"16px"},"body-style":{padding:"0px 16px"}},{extra:r(()=>[t.open!==void 0?(a(),n(s,{key:0,modelValue:t.open,"onUpdate:modelValue":S=>t.open=S,onChange:S=>K(t)},null,8,["modelValue","onUpdate:modelValue","onChange"])):_("",!0)]),actions:r(()=>[t.reset?(a(),n(O,{key:0,onClick:S=>L(t)},{default:r(()=>[B(U(u.$t("button.reset")),1)]),_:2},1032,["onClick"])):_("",!0),t.config?(a(),n(O,{key:1,type:"primary",onClick:S=>H(t)},{default:r(()=>[B(U(u.$t("button.config")),1)]),_:2},1032,["onClick"])):_("",!0)]),skeleton:r(()=>[l(T,{animation:!0},{default:r(()=>[l(R,{widths:["50%","50%","100%","40%"],rows:4}),l(R,{widths:["40%"],rows:1})]),_:1})]),default:r(()=>[l($,null,{description:r(()=>[B(U(t.description)+" ",1),l(g,{style:{"margin-top":"10px"},data:t.data,layout:"inline-horizontal",column:2},null,8,["data"])]),_:2},1024)]),_:2},1032,["title"])])]),_:2},1024))),128))]),_:1}),l(j,{visible:v.value,"onUpdate:visible":c[3]||(c[3]=t=>v.value=t),title:u.$t(q.value),onCancel:x,onBeforeOk:G},{default:r(()=>[l(I,{ref_key:"configForm",ref:k,model:e.value,"auto-label-width":""},{default:r(()=>[e.value.action==="statistics"?(a(),n(D,{key:0,field:"statistics.cron",label:u.$t("sys.config.label.statistics.cron"),rules:[{required:!0,message:u.$t("sys.config.error.statistics.cron.required")}]},{default:r(()=>[l(A,{modelValue:e.value.statistics.cron,"onUpdate:modelValue":c[0]||(c[0]=t=>e.value.statistics.cron=t),placeholder:u.$t("sys.config.placeholder.statistics.cron"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"])):_("",!0),e.value.action==="statistics"?(a(),n(D,{key:1,field:"statistics.limit",label:u.$t("sys.config.label.statistics.limit"),rules:[{required:!0,message:u.$t("sys.config.error.statistics.limit.required")}]},{default:r(()=>[l(F,{modelValue:e.value.statistics.limit,"onUpdate:modelValue":c[1]||(c[1]=t=>e.value.statistics.limit=t),placeholder:u.$t("sys.config.placeholder.statistics.limit"),min:10,"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"])):_("",!0),e.value.action==="statistics"?(a(),n(D,{key:2,field:"statistics.lock_minutes",label:u.$t("sys.config.label.statistics.lock_minutes"),rules:[{required:!0,message:u.$t("sys.config.error.statistics.lock_minutes.required")}]},{default:r(()=>[l(F,{modelValue:e.value.statistics.lock_minutes,"onUpdate:modelValue":c[2]||(c[2]=t=>e.value.statistics.lock_minutes=t),placeholder:u.$t("sys.config.placeholder.statistics.lock_minutes"),min:1,"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"])):_("",!0)]),_:1},8,["model"])]),_:1},8,["visible","title"])])}}});const we=ee(el,[["__scopeId","data-v-737f5686"]]),ll={class:"list-wrap",style:{"margin-top":"10px"}},rl={class:"card-wrap"},ul={name:"API"},ol=le({...ul,setup(ae){const{proxy:b}=Ee(),{setLoading:f}=he(!0),{t:y}=ne(),v=E(!1),q=E(""),k=E(),e=E({base:{},log:{},auto_disabled_error:{},auto_enable_error:{},not_retry_error:{},not_shield_error:{}}),H=async o=>{o.action==="auto_disabled_error"&&e.value.auto_disabled_error.errors.length===0&&u(),o.action==="auto_enable_error"&&e.value.auto_enable_error.enable_errors.length===0&&s(),o.action==="not_retry_error"&&e.value.not_retry_error.errors.length===0&&$(),o.action==="not_shield_error"&&e.value.not_shield_error.errors.length===0&&R(),q.value=y(`sys.config.item.title.${o.action}`),e.value.action=o.action,v.value=!0},L=async o=>{b.$modal.warning({title:"\u8B66\u544A",titleAlign:"center",content:`\u662F\u5426\u786E\u5B9A\u91CD\u7F6E${y(`sys.config.item.title.${o.action}`)}?`,hideCancel:!1,onOk:()=>{J(o)}})},G=async o=>{var z;if(await((z=k.value)==null?void 0:z.validate())){v.value=!0,o(!1);return}f(!0);try{await Ce(e.value),o(),b.$message.success("\u64CD\u4F5C\u6210\u529F"),h()}catch{o(!1)}finally{f(!1)}},x=()=>{v.value=!1,e.value.auto_disabled_error.errors.length>0&&!e.value.auto_disabled_error.errors[e.value.auto_disabled_error.errors.length-1]&&c(e.value.auto_disabled_error.errors.length-1),e.value.auto_enable_error.enable_errors.length>0&&(!e.value.auto_enable_error.enable_errors[e.value.auto_enable_error.enable_errors.length-1].enable_time||!e.value.auto_enable_error.enable_errors[e.value.auto_enable_error.enable_errors.length-1].cron||!e.value.auto_enable_error.enable_errors[e.value.auto_enable_error.enable_errors.length-1].error)&&g(e.value.auto_enable_error.enable_errors.length-1),e.value.not_retry_error.errors.length>0&&!e.value.not_retry_error.errors[e.value.not_retry_error.errors.length-1]&&O(e.value.not_retry_error.errors.length-1),e.value.not_shield_error.errors.length>0&&!e.value.not_shield_error.errors[e.value.not_shield_error.errors.length-1]&&T(e.value.not_shield_error.errors.length-1)},J=async o=>{f(!0);try{await $e({action:o.action}),b.$message.success("\u64CD\u4F5C\u6210\u529F"),h()}finally{f(!1)}},K=async o=>{f(!0);try{await Be({action:o.action,open:o.open||!1}),b.$message.success("\u64CD\u4F5C\u6210\u529F"),h()}finally{f(!1)}},N=E({}),h=async()=>{const{data:o}=await ke();e.value.base=o.base,e.value.log=o.log,e.value.auto_disabled_error=o.auto_disabled_error,e.value.auto_enable_error=o.auto_enable_error,e.value.not_retry_error=o.not_retry_error,e.value.not_shield_error=o.not_shield_error,N.value=[{action:"base",title:y("sys.config.item.title.base"),description:"\u914D\u7F6E\u9519\u8BEF\u91CD\u8BD5\u6B21\u6570\u548C\u5404\u7C7B\u9519\u8BEF\u7981\u7528\u6B21\u6570\u7B49, \u9519\u8BEF\u91CD\u8BD5\u6B21\u6570N > 0 \u91CD\u8BD5 N \u6B21, N < 0 \u91CD\u8BD5\u6240\u6709key\u4E00\u8F6E, N = 0 \u4E0D\u91CD\u8BD5, \u9519\u8BEF\u6B21\u6570\u6BCF\u59290\u70B9\u4F1A\u81EA\u52A8\u91CD\u7F6E, \u6CE8: \u4EE3\u7406\u5BC6\u94A5\u9519\u8BEF\u65F6, \u4E5F\u4F1A\u8BB0\u5F55\u6A21\u578B\u4EE3\u7406\u9519\u8BEF\u6B21\u6570",config:!0,reset:!0},{action:"log",title:y("sys.config.item.title.log"),description:"\u8C03\u7528\u65E5\u5FD7\u8BB0\u5F55\u5185\u5BB9, \u652F\u6301\u8BB0\u5F55: \u63D0\u95EE\u3001\u56DE\u7B54\u3001\u4E0A\u4E0B\u6587\u3001\u591A\u6A21\u6001\u8BC6\u56FE\u7684BASE64\u56FE\u50CF\u6570\u636E",open:e.value.log.open,config:!0,reset:!0},{action:"auto_disabled_error",title:y("sys.config.item.title.auto_disabled_error"),description:"\u8C03\u7528\u62A5\u9519\u65F6, \u5305\u542B\u6709\u914D\u7F6E\u9519\u8BEF\u5185\u5BB9\u65F6\u5219\u81EA\u52A8\u4F1A\u7981\u7528\u5BC6\u94A5\u6216\u6A21\u578B\u4EE3\u7406\u7B49, \u4E3A\u7A7A\u5219\u4E0D\u4F1A\u81EA\u52A8\u7981\u7528(\u8FBE\u5230\u9519\u8BEF\u6B21\u6570\u4E0A\u9650\u9664\u5916)",open:e.value.auto_disabled_error.open,config:!0,reset:!0},{action:"auto_enable_error",title:y("sys.config.item.title.auto_enable_error"),description:"\u5BC6\u94A5\u81EA\u52A8\u7981\u7528\u540E, \u53EF\u901A\u8FC7\u6B64\u914D\u7F6E\u81EA\u52A8\u542F\u7528, \u5BC6\u94A5\u7981\u7528\u539F\u56E0\u5305\u542B\u6709\u914D\u7F6E\u7684\u9519\u8BEF\u5185\u5BB9\u65F6, \u4F1A\u6839\u636E\u914D\u7F6E\u7684\u542F\u7528\u65F6\u95F4\u5224\u65AD\u662F\u5426\u6EE1\u8DB3\u542F\u7528\u6761\u4EF6, \u6EE1\u8DB3\u5219\u4F1A\u81EA\u52A8\u542F\u7528, \u542F\u7528\u65F6\u95F4\u5355\u4F4D: \u79D2",open:e.value.auto_enable_error.open,config:!0,reset:!0},{action:"not_retry_error",title:y("sys.config.item.title.not_retry_error"),description:"\u8C03\u7528\u62A5\u9519\u65F6, \u5305\u542B\u6709\u914D\u7F6E\u9519\u8BEF\u5185\u5BB9\u65F6\u5219\u4E0D\u4F1A\u81EA\u52A8\u91CD\u8BD5, \u4E3A\u7A7A\u5219\u4F1A\u81EA\u52A8\u91CD\u8BD5",open:e.value.not_retry_error.open,config:!0,reset:!0},{action:"not_shield_error",title:y("sys.config.item.title.not_shield_error"),description:"\u8C03\u7528\u62A5\u9519\u65F6, \u5305\u542B\u6709\u914D\u7F6E\u9519\u8BEF\u5185\u5BB9\u65F6\u5219\u4F1A\u5C06\u9519\u8BEF\u5185\u5BB9\u8FD4\u56DE\u7ED9\u8C03\u7528\u65B9, \u4E3A\u7A7A\u5219\u5C4F\u853D\u6240\u6709\u9519\u8BEF",open:e.value.not_shield_error.open,config:!0,reset:!0},{action:"reset_api_error",title:y("sys.config.item.title.reset_api_error"),description:"\u5F53\u9519\u8BEF\u6B21\u6570\u8FBE\u5230\u914D\u7F6E\u4E0A\u9650\u65F6\u53EF\u624B\u52A8\u8FDB\u884C\u91CD\u7F6E, \u91CD\u7F6E\u8FC7\u7A0B\u53EF\u80FD\u4F1A\u9020\u6210\u7CFB\u7EDF\u7684\u77ED\u6682\u4E0D\u53EF\u7528(\u4E00\u822C\u51E0\u79D2\u949F), \u8BF7\u8C28\u614E\u64CD\u4F5C, \u6216\u5C1D\u8BD5\u8C03\u9AD8\u57FA\u7840\u914D\u7F6E\u7684\u9519\u8BEF\u6B21\u6570, \u8C03\u9AD8\u4E0D\u4F1A\u5F71\u54CD\u7CFB\u7EDF\u7684\u6B63\u5E38\u8FD0\u884C",reset:!0}]};h();const u=()=>{e.value.auto_disabled_error.errors.push("")},c=o=>{e.value.auto_disabled_error.errors.splice(o,1)},s=()=>{e.value.auto_enable_error.enable_errors.push({cron:"",enable_time:E(),error:""})},g=o=>{e.value.log.records.length>1&&e.value.auto_enable_error.enable_errors.splice(o,1)},$=()=>{e.value.not_retry_error.errors.push("")},O=o=>{e.value.not_retry_error.errors.splice(o,1)},R=()=>{e.value.not_shield_error.errors.push("")},T=o=>{e.value.not_shield_error.errors.splice(o,1)};return(o,m)=>{const z=de,A=_e,D=ce,F=pe,I=me,j=ge,t=re,S=ue,De=oe,Q=be,i=ye,V=He,W=Le,X=fe,te=qe,se=Se,Ne=ve,Oe=Fe;return a(),w("div",ll,[l(De,{class:"list-row",gutter:24},{default:r(()=>[(a(!0),w(P,null,M(N.value,p=>(a(),n(S,{key:p.action,xs:12,sm:12,md:12,lg:6,xl:6,xxl:6,class:"list-col"},{default:r(()=>[ie("div",rl,[l(t,{title:p.title,bordered:!1,"header-style":{padding:"16px"},"body-style":{padding:"0px 16px"}},{extra:r(()=>[p.open!==void 0?(a(),n(z,{key:0,modelValue:p.open,"onUpdate:modelValue":d=>p.open=d,onChange:d=>K(p)},null,8,["modelValue","onUpdate:modelValue","onChange"])):_("",!0)]),actions:r(()=>[p.reset?(a(),n(F,{key:0,onClick:d=>L(p)},{default:r(()=>[B(U(o.$t("button.reset")),1)]),_:2},1032,["onClick"])):_("",!0),p.config?(a(),n(F,{key:1,type:"primary",onClick:d=>H(p)},{default:r(()=>[B(U(o.$t("button.config")),1)]),_:2},1032,["onClick"])):_("",!0)]),skeleton:r(()=>[l(j,{animation:!0},{default:r(()=>[l(I,{widths:["50%","50%","100%","40%"],rows:4}),l(I,{widths:["40%"],rows:1})]),_:1})]),default:r(()=>[l(D,null,{description:r(()=>[B(U(p.description)+" ",1),l(A,{style:{"margin-top":"10px"},data:p.data,layout:"inline-horizontal",column:2},null,8,["data"])]),_:2},1024)]),_:2},1032,["title"])])]),_:2},1024))),128))]),_:1}),l(Oe,{visible:v.value,"onUpdate:visible":m[12]||(m[12]=p=>v.value=p),title:o.$t(q.value),width:e.value.action==="base"||e.value.action==="log"?520:700,onCancel:x,onBeforeOk:G},{default:r(()=>[l(Ne,{ref_key:"configForm",ref:k,model:e.value,"auto-label-width":"",style:{"max-height":"300px"}},{default:r(()=>[e.value.action==="base"?(a(),n(i,{key:0,field:"base.err_retry",label:o.$t("sys.config.label.base.err_retry"),rules:[{required:!0,message:o.$t("sys.config.error.base.err_retry.required")}]},{default:r(()=>[l(Q,{modelValue:e.value.base.err_retry,"onUpdate:modelValue":m[0]||(m[0]=p=>e.value.base.err_retry=p),placeholder:o.$t("sys.config.placeholder.base.err_retry"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"])):_("",!0),e.value.action==="base"?(a(),n(i,{key:1,field:"base.model_key_err_disable",label:o.$t("sys.config.label.base.model_key_err_disable"),rules:[{required:!0,message:o.$t("sys.config.error.base.model_key_err_disable.required")}]},{default:r(()=>[l(Q,{modelValue:e.value.base.model_key_err_disable,"onUpdate:modelValue":m[1]||(m[1]=p=>e.value.base.model_key_err_disable=p),placeholder:o.$t("sys.config.placeholder.base.model_key_err_disable"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"])):_("",!0),e.value.action==="base"?(a(),n(i,{key:2,field:"base.model_agent_err_disable",label:o.$t("sys.config.label.base.model_agent_err_disable"),rules:[{required:!0,message:o.$t("sys.config.error.base.model_agent_err_disable.required")}]},{default:r(()=>[l(Q,{modelValue:e.value.base.model_agent_err_disable,"onUpdate:modelValue":m[2]||(m[2]=p=>e.value.base.model_agent_err_disable=p),placeholder:o.$t("sys.config.placeholder.base.model_agent_err_disable"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"])):_("",!0),e.value.action==="base"?(a(),n(i,{key:3,field:"base.model_agent_key_err_disable",label:o.$t("sys.config.label.base.model_agent_key_err_disable"),rules:[{required:!0,message:o.$t("sys.config.error.base.model_agent_key_err_disable.required")}]},{default:r(()=>[l(Q,{modelValue:e.value.base.model_agent_key_err_disable,"onUpdate:modelValue":m[3]||(m[3]=p=>e.value.base.model_agent_key_err_disable=p),placeholder:o.$t("sys.config.placeholder.base.model_agent_key_err_disable"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"])):_("",!0),e.value.action==="log"?(a(),n(i,{key:4,field:"log.records",label:o.$t("sys.config.label.log.records")},{default:r(()=>[l(W,{size:"large"},{default:r(()=>[l(V,{modelValue:e.value.log.records,"onUpdate:modelValue":m[4]||(m[4]=p=>e.value.log.records=p),value:"prompt"},{default:r(()=>[B(" \u63D0\u95EE ")]),_:1},8,["modelValue"]),l(V,{modelValue:e.value.log.records,"onUpdate:modelValue":m[5]||(m[5]=p=>e.value.log.records=p),value:"completion"},{default:r(()=>[B(" \u56DE\u7B54 ")]),_:1},8,["modelValue"]),l(V,{modelValue:e.value.log.records,"onUpdate:modelValue":m[6]||(m[6]=p=>e.value.log.records=p),value:"messages"},{default:r(()=>[B(" \u4E0A\u4E0B\u6587 ")]),_:1},8,["modelValue"]),l(V,{modelValue:e.value.log.records,"onUpdate:modelValue":m[7]||(m[7]=p=>e.value.log.records=p),value:"image"},{default:r(()=>[B(" \u8BC6\u56FE\u7684\u56FE\u50CF\u6570\u636E ")]),_:1},8,["modelValue"])]),_:1})]),_:1},8,["label"])):_("",!0),(a(!0),w(P,null,M(e.value.auto_disabled_error.errors,(p,d)=>Y((a(),n(i,{key:d,field:`auto_disabled_error.errors[${d}]`,label:`${d+1}. `+o.$t("sys.config.label.auto_disabled_error.errors"),rules:[{required:!0,message:o.$t("sys.config.error.auto_disabled_error.errors.required")}],"label-col-style":{padding:"0 16px 2px 0"}},{default:r(()=>[l(X,{modelValue:e.value.auto_disabled_error.errors[d],"onUpdate:modelValue":C=>e.value.auto_disabled_error.errors[d]=C,placeholder:o.$t("sys.config.placeholder.auto_disabled_error.errors"),"allow-clear":"",style:{width:"85%","margin-right":"5px"}},null,8,["modelValue","onUpdate:modelValue","placeholder"]),l(F,{type:"primary",shape:"circle",style:{margin:"0 10px 0 10px"},onClick:m[8]||(m[8]=C=>u())},{default:r(()=>[l(te)]),_:1}),l(F,{type:"secondary",shape:"circle",onClick:C=>c(d)},{default:r(()=>[l(se)]),_:2},1032,["onClick"])]),_:2},1032,["field","label","rules"])),[[Z,e.value.action==="auto_disabled_error"]])),128)),(a(!0),w(P,null,M(e.value.auto_enable_error.enable_errors,(p,d)=>Y((a(),n(i,{key:d,field:`auto_enable_error.enable_errors[${d}].cron`&&`auto_enable_error.enable_errors[${d}].enable_time`&&`auto_enable_error.enable_errors[${d}].error`,label:`${d+1}.`,rules:[{required:!0,message:o.$t("sys.config.error.auto_enable_error.required")}],"label-col-style":{padding:"0 16px 2px 0"}},{default:r(()=>[l(X,{modelValue:e.value.auto_enable_error.enable_errors[d].cron,"onUpdate:modelValue":C=>e.value.auto_enable_error.enable_errors[d].cron=C,placeholder:o.$t("sys.config.placeholder.auto_enable_error.enable_errors.cron"),"allow-clear":"",style:{width:"122px","margin-right":"5px"}},null,8,["modelValue","onUpdate:modelValue","placeholder"]),l(Q,{modelValue:e.value.auto_enable_error.enable_errors[d].enable_time,"onUpdate:modelValue":C=>e.value.auto_enable_error.enable_errors[d].enable_time=C,placeholder:o.$t("sys.config.placeholder.auto_enable_error.enable_errors.enable_time"),"allow-clear":"",style:{width:"95px","margin-right":"5px"}},null,8,["modelValue","onUpdate:modelValue","placeholder"]),l(X,{modelValue:e.value.auto_enable_error.enable_errors[d].error,"onUpdate:modelValue":C=>e.value.auto_enable_error.enable_errors[d].error=C,placeholder:o.$t("sys.config.placeholder.auto_enable_error.enable_errors.error"),"allow-clear":"",style:{width:"305px","margin-right":"5px"}},null,8,["modelValue","onUpdate:modelValue","placeholder"]),l(F,{type:"primary",shape:"circle",style:{margin:"0 10px 0 10px"},onClick:m[9]||(m[9]=C=>s())},{default:r(()=>[l(te)]),_:1}),l(F,{type:"secondary",shape:"circle",onClick:C=>g(d)},{default:r(()=>[l(se)]),_:2},1032,["onClick"])]),_:2},1032,["field","label","rules"])),[[Z,e.value.action==="auto_enable_error"]])),128)),(a(!0),w(P,null,M(e.value.not_retry_error.errors,(p,d)=>Y((a(),n(i,{key:d,field:`not_retry_error.errors[${d}]`,label:`${d+1}. `+o.$t("sys.config.label.not_retry_error.errors"),rules:[{required:!0,message:o.$t("sys.config.error.not_retry_error.errors.required")}],"label-col-style":{padding:"0 16px 2px 0"}},{default:r(()=>[l(X,{modelValue:e.value.not_retry_error.errors[d],"onUpdate:modelValue":C=>e.value.not_retry_error.errors[d]=C,placeholder:o.$t("sys.config.placeholder.not_retry_error.errors"),"allow-clear":"",style:{width:"85%","margin-right":"5px"}},null,8,["modelValue","onUpdate:modelValue","placeholder"]),l(F,{type:"primary",shape:"circle",style:{margin:"0 10px 0 10px"},onClick:m[10]||(m[10]=C=>$())},{default:r(()=>[l(te)]),_:1}),l(F,{type:"secondary",shape:"circle",onClick:C=>O(d)},{default:r(()=>[l(se)]),_:2},1032,["onClick"])]),_:2},1032,["field","label","rules"])),[[Z,e.value.action==="not_retry_error"]])),128)),(a(!0),w(P,null,M(e.value.not_shield_error.errors,(p,d)=>Y((a(),n(i,{key:d,field:`not_shield_error.errors[${d}]`,label:`${d+1}. `+o.$t("sys.config.label.not_shield_error.errors"),rules:[{required:!0,message:o.$t("sys.config.error.not_shield_error.errors.required")}],"label-col-style":{padding:"0 16px 2px 0"}},{default:r(()=>[l(X,{modelValue:e.value.not_shield_error.errors[d],"onUpdate:modelValue":C=>e.value.not_shield_error.errors[d]=C,placeholder:o.$t("sys.config.placeholder.not_shield_error.errors"),"allow-clear":"",style:{width:"85%","margin-right":"5px"}},null,8,["modelValue","onUpdate:modelValue","placeholder"]),l(F,{type:"primary",shape:"circle",style:{margin:"0 10px 0 10px"},onClick:m[11]||(m[11]=C=>R())},{default:r(()=>[l(te)]),_:1}),l(F,{type:"secondary",shape:"circle",onClick:C=>T(d)},{default:r(()=>[l(se)]),_:2},1032,["onClick"])]),_:2},1032,["field","label","rules"])),[[Z,e.value.action==="not_shield_error"]])),128))]),_:1},8,["model"])]),_:1},8,["visible","title","width"])])}}});const Ue=ee(ol,[["__scopeId","data-v-2c0d7b9c"]]),al={class:"container"},tl={name:"SysConfig"},sl=le({...tl,setup(ae){return(b,f)=>{const y=Te,v=Re,q=ze,k=Ge,e=Ie,H=ue,L=oe,G=re;return a(),w("div",al,[l(q,{class:"container-breadcrumb"},{default:r(()=>[l(v,null,{default:r(()=>[l(y)]),_:1}),l(v,null,{default:r(()=>[B(U(b.$t("menu.sys")),1)]),_:1}),l(v,null,{default:r(()=>[B(U(b.$t("menu.sys.config")),1)]),_:1})]),_:1}),l(L,{gutter:20,align:"stretch"},{default:r(()=>[l(H,{span:24},{default:r(()=>[l(G,{class:"general-card",bordered:!1,"header-style":{padding:"20px"},"body-style":{padding:"25px 20px 20px 14px"}},{default:r(()=>[l(L,{justify:"space-between"},{default:r(()=>[l(H,{span:24},{default:r(()=>[l(e,{"default-active-tab":0,type:"rounded"},{default:r(()=>[l(k,{key:"0",title:b.$t("sys.config.tab.title.all")},{default:r(()=>[l(Ve),l(Ae),l(we),l(Ue)]),_:1},8,["title"]),l(k,{key:"1",title:b.$t("sys.config.tab.title.global")},{default:r(()=>[l(Ve)]),_:1},8,["title"]),l(k,{key:"2",title:b.$t("sys.config.tab.title.general")},{default:r(()=>[l(Ae)]),_:1},8,["title"]),l(k,{key:"3",title:b.$t("sys.config.tab.title.task")},{default:r(()=>[l(we)]),_:1},8,["title"]),l(k,{key:"4",title:b.$t("sys.config.tab.title.api")},{default:r(()=>[l(Ue)]),_:1},8,["title"])]),_:1})]),_:1})]),_:1})]),_:1})]),_:1})]),_:1})])}}});const Cl=ee(sl,[["__scopeId","data-v-2861651e"]]);export{Cl as default};
