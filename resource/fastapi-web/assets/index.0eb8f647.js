import{c as se,u as ne,_ as ee,i as Te}from"./index.47a80abf.js";/* empty css                *//* empty css                *//* empty css               *//* empty css               */import{d as le,e as v,B as u,C as V,aH as l,aG as t,aJ as z,aI as M,aD as n,F as ie,aE as _,aL as A,aM as U,aT as ce,bN as de,c3 as _e,aU as pe,bJ as me,bK as fe,bI as te,bC as oe,bD as re,b1 as ye,b2 as ge,aS as De,b5 as ve,a_ as be,g as he,bu as Y,bv as Z,a5 as qe,a4 as Se,aK as He,aF as Le,bb as Re,b6 as xe}from"./arco.a9260898.js";/* empty css               *//* empty css                *//* empty css                */import{u as Fe}from"./loading.1f346a94.js";import"./chart.d103b168.js";import"./vue.143177a1.js";function $e(){return se.get("/api/v1/sys/config/detail")}function Ce(L){return se.post("/api/v1/sys/config/update",L)}function Ee(L){return se.post("/api/v1/sys/config/change/status",L)}function ke(L){return se.post("/api/v1/sys/config/reset",L)}const Ie={class:"list-wrap",style:{"margin-top":"10px"}},Ge={class:"card-wrap"},Pe={name:"Global"},ze=le({...Pe,setup(L){const{proxy:p}=he(),{setLoading:d}=Fe(!0),{t:y}=ne(),m=v(!1),q=v(""),C=v(),e=v({email:{},http:{},core:{}}),R=async o=>{q.value=y(`sys.config.item.title.${o.action}`),e.value.action=o.action,m.value=!0},x=async o=>{p.$modal.warning({title:"\u8B66\u544A",titleAlign:"center",content:`\u662F\u5426\u786E\u5B9A\u91CD\u7F6E${y(`sys.config.item.title.${o.action}`)}?`,hideCancel:!1,onOk:()=>{Q(o)}})},P=async o=>{var s;if(await((s=C.value)==null?void 0:s.validate())){m.value=!0,o(!1);return}d(!0);try{await Ce(e.value),o(),p.$message.success("\u64CD\u4F5C\u6210\u529F"),F()}catch{o(!1)}finally{d(!1)}},j=()=>{m.value=!1},Q=async o=>{d(!0);try{await ke({action:o.action}),p.$message.success("\u64CD\u4F5C\u6210\u529F"),F()}finally{d(!1)}},W=async o=>{d(!0);try{await Ee({action:o.action,open:o.open||!1}),p.$message.success("\u64CD\u4F5C\u6210\u529F"),F()}finally{d(!1)}},k=v({}),N=v({}),F=async()=>{const{data:o}=await $e();k.value=o,e.value.email=o.email,e.value.http=o.http,e.value.core=o.core,N.value=[{action:"email",title:y("sys.config.item.title.email"),description:"\u914D\u7F6E\u53D1\u4FE1\u90AE\u7BB1\u4FE1\u606F, \u5F53\u6709\u7AD9\u70B9\u914D\u7F6E\u4E14\u540C\u65F6\u914D\u7F6E\u53D1\u4FE1\u90AE\u7BB1\u4FE1\u606F\u65F6, \u7AD9\u70B9\u914D\u7F6E\u7684\u4F18\u5148",open:k.value.email.open,config:!0,reset:!0},{action:"http",title:y("sys.config.item.title.http"),description:"\u914D\u7F6EHTTP\u8BF7\u6C42\u8D85\u65F6\u65F6\u95F4\u548C\u4EE3\u7406\u5730\u5740",config:!0,reset:!0},{action:"core",title:y("sys.config.item.title.core"),description:"\u7CFB\u7EDF\u9996\u6B21\u4F7F\u7528\u65F6\u53EF\u6839\u636E\u5B9E\u9645\u60C5\u51B5\u4FEE\u6539, \u540E\u671F\u5982\u82E5\u518D\u6B21\u4FEE\u6539\u5BC6\u94A5\u524D\u7F00, \u5386\u53F2\u5E94\u7528\u5BC6\u94A5\u5C06\u65E0\u6CD5\u901A\u8FC7\u6838\u9A8C, \u8BF7\u8C28\u614E\u4FEE\u6539",config:!0,reset:!0},{action:"debug",title:y("sys.config.item.title.debug"),description:"\u8C03\u8BD5\u5F00\u5173\u6253\u5F00\u540E, \u65E5\u5FD7\u4F1A\u6253\u5370\u66F4\u591A\u8BE6\u7EC6\u4FE1\u606F, \u65E5\u5FD7\u7EA7\u522B(logger.level)\u9700\u914D\u7F6E\u4E3A: debug",open:k.value.debug.open}]};return F(),(o,i)=>{const s=ce,D=de,O=_e,T=pe,I=me,H=fe,r=te,g=oe,G=re,w=ye,E=ge,$=De,J=ve,X=be;return u(),V("div",Ie,[l(G,{class:"list-row",gutter:24},{default:t(()=>[(u(!0),V(z,null,M(N.value,a=>(u(),n(g,{key:a.action,xs:12,sm:12,md:12,lg:6,xl:6,xxl:6,class:"list-col"},{default:t(()=>[ie("div",Ge,[l(r,{title:a.title,bordered:!1,"header-style":{padding:"16px"},"body-style":{padding:"0px 16px"}},{extra:t(()=>[a.open!==void 0?(u(),n(s,{key:0,modelValue:a.open,"onUpdate:modelValue":S=>a.open=S,onChange:S=>W(a)},null,8,["modelValue","onUpdate:modelValue","onChange"])):_("",!0)]),actions:t(()=>[a.reset?(u(),n(T,{key:0,onClick:S=>x(a)},{default:t(()=>[A(U(o.$t("button.reset")),1)]),_:2},1032,["onClick"])):_("",!0),a.config?(u(),n(T,{key:1,type:"primary",onClick:S=>R(a)},{default:t(()=>[A(U(o.$t("button.config")),1)]),_:2},1032,["onClick"])):_("",!0)]),skeleton:t(()=>[l(H,{animation:!0},{default:t(()=>[l(I,{widths:["50%","50%","100%","40%"],rows:4}),l(I,{widths:["40%"],rows:1})]),_:1})]),default:t(()=>[l(O,null,{description:t(()=>[A(U(a.description)+" ",1),l(D,{style:{"margin-top":"10px"},data:a.data,layout:"inline-horizontal",column:2},null,8,["data"])]),_:2},1024)]),_:2},1032,["title"])])]),_:2},1024))),128))]),_:1}),l(X,{visible:m.value,"onUpdate:visible":i[9]||(i[9]=a=>m.value=a),title:o.$t(q.value),onCancel:j,onBeforeOk:P},{default:t(()=>[l(J,{ref_key:"configForm",ref:C,model:e.value,"auto-label-width":""},{default:t(()=>[e.value.action==="email"?(u(),n(E,{key:0,field:"email.host",label:o.$t("sys.config.label.email.host"),rules:[{required:!0,message:o.$t("sys.config.error.email.host.required")}]},{default:t(()=>[l(w,{modelValue:e.value.email.host,"onUpdate:modelValue":i[0]||(i[0]=a=>e.value.email.host=a),placeholder:o.$t("sys.config.placeholder.email.host"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"])):_("",!0),e.value.action==="email"?(u(),n(E,{key:1,field:"email.port",label:o.$t("sys.config.label.email.port"),rules:[{required:!0,message:o.$t("sys.config.error.email.port.required")}]},{default:t(()=>[l($,{modelValue:e.value.email.port,"onUpdate:modelValue":i[1]||(i[1]=a=>e.value.email.port=a),placeholder:o.$t("sys.config.placeholder.email.port"),min:1,"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"])):_("",!0),e.value.action==="email"?(u(),n(E,{key:2,field:"email.user_name",label:o.$t("sys.config.label.email.user_name"),rules:[{required:!0,message:o.$t("sys.config.error.email.user_name.required")}]},{default:t(()=>[l(w,{modelValue:e.value.email.user_name,"onUpdate:modelValue":i[2]||(i[2]=a=>e.value.email.user_name=a),placeholder:o.$t("sys.config.placeholder.email.user_name"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"])):_("",!0),e.value.action==="email"?(u(),n(E,{key:3,field:"email.password",label:o.$t("sys.config.label.email.password"),rules:[{required:!0,message:o.$t("sys.config.error.email.password.required")}]},{default:t(()=>[l(w,{modelValue:e.value.email.password,"onUpdate:modelValue":i[3]||(i[3]=a=>e.value.email.password=a),placeholder:o.$t("sys.config.placeholder.email.password"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"])):_("",!0),e.value.action==="email"?(u(),n(E,{key:4,field:"email.from_name",label:o.$t("sys.config.label.email.from_name"),rules:[{required:!0,message:o.$t("sys.config.error.email.from_name.required")}]},{default:t(()=>[l(w,{modelValue:e.value.email.from_name,"onUpdate:modelValue":i[4]||(i[4]=a=>e.value.email.from_name=a),placeholder:o.$t("sys.config.placeholder.email.from_name"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"])):_("",!0),e.value.action==="http"?(u(),n(E,{key:5,field:"http.timeout",label:o.$t("sys.config.label.http.timeout"),rules:[{required:!0,message:o.$t("sys.config.error.http.timeout.required")}]},{default:t(()=>[l($,{modelValue:e.value.http.timeout,"onUpdate:modelValue":i[5]||(i[5]=a=>e.value.http.timeout=a),placeholder:o.$t("sys.config.placeholder.http.timeout"),min:1,"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"])):_("",!0),e.value.action==="http"?(u(),n(E,{key:6,field:"http.proxy_url",label:o.$t("sys.config.label.http.proxy_url")},{default:t(()=>[l(w,{modelValue:e.value.http.proxy_url,"onUpdate:modelValue":i[6]||(i[6]=a=>e.value.http.proxy_url=a),placeholder:o.$t("sys.config.placeholder.http.proxy_url"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label"])):_("",!0),e.value.action==="core"?(u(),n(E,{key:7,field:"core.secret_key_prefix",label:o.$t("sys.config.label.core.secret_key_prefix")},{default:t(()=>[l(w,{modelValue:e.value.core.secret_key_prefix,"onUpdate:modelValue":i[7]||(i[7]=a=>e.value.core.secret_key_prefix=a),placeholder:o.$t("sys.config.placeholder.core.secret_key_prefix"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label"])):_("",!0),e.value.action==="core"?(u(),n(E,{key:8,field:"core.error_prefix",label:o.$t("sys.config.label.core.error_prefix")},{default:t(()=>[l(w,{modelValue:e.value.core.error_prefix,"onUpdate:modelValue":i[8]||(i[8]=a=>e.value.core.error_prefix=a),placeholder:o.$t("sys.config.placeholder.core.error_prefix"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label"])):_("",!0)]),_:1},8,["model"])]),_:1},8,["visible","title"])])}}});const we=ee(ze,[["__scopeId","data-v-9f920a93"]]),Me={class:"list-wrap",style:{"margin-top":"10px"}},Je={class:"card-wrap"},Ke={name:"General"},je=le({...Ke,setup(L){const{proxy:p}=he(),{setLoading:d}=Fe(!0),{t:y}=ne(),m=v(!1),q=v(""),C=v(),e=v({user_shield_error:{}}),R=async s=>{s.action==="user_shield_error"&&e.value.user_shield_error.errors.length===0&&o(),q.value=y(`sys.config.item.title.${s.action}`),e.value.action=s.action,m.value=!0},x=async s=>{p.$modal.warning({title:"\u8B66\u544A",titleAlign:"center",content:`\u662F\u5426\u786E\u5B9A\u91CD\u7F6E${y(`sys.config.item.title.${s.action}`)}?`,hideCancel:!1,onOk:()=>{Q(s)}})},P=async s=>{var O;if(await((O=C.value)==null?void 0:O.validate())){m.value=!0,s(!1);return}d(!0);try{await Ce(e.value),s(),p.$message.success("\u64CD\u4F5C\u6210\u529F"),F()}catch{s(!1)}finally{d(!1)}},j=()=>{m.value=!1,e.value.user_shield_error.errors.length>0&&!e.value.user_shield_error.errors[e.value.user_shield_error.errors.length-1]&&i(e.value.user_shield_error.errors.length-1)},Q=async s=>{d(!0);try{await ke({action:s.action}),p.$message.success("\u64CD\u4F5C\u6210\u529F"),F()}finally{d(!1)}},W=async s=>{d(!0);try{await Ee({action:s.action,open:s.open||!1}),p.$message.success("\u64CD\u4F5C\u6210\u529F"),F()}finally{d(!1)}},k=v({}),N=v({}),F=async()=>{const{data:s}=await $e();k.value=s,e.value.user_shield_error=s.user_shield_error,N.value=[{action:"user_shield_error",title:y("sys.config.item.title.user_shield_error"),description:"\u7528\u6237\u67E5\u770B\u8C03\u7528\u65E5\u5FD7\u9519\u8BEF\u65F6, \u5305\u542B\u6709\u914D\u7F6E\u9519\u8BEF\u5185\u5BB9\u65F6\u5219\u5C4F\u853D\u663E\u793A, \u4E3A\u7A7A\u5219\u5C4F\u853D\u6240\u6709\u9519\u8BEF\u663E\u793A",open:k.value.user_shield_error.open,config:!0,reset:!0}]};F();const o=()=>{e.value.user_shield_error.errors.push("")},i=s=>{e.value.user_shield_error.errors.splice(s,1)};return(s,D)=>{const O=ce,T=de,I=_e,H=pe,r=me,g=fe,G=te,w=oe,E=re,$=ye,J=qe,X=Se,a=ge,S=ve,Be=be;return u(),V("div",Me,[l(E,{class:"list-row",gutter:24},{default:t(()=>[(u(!0),V(z,null,M(N.value,b=>(u(),n(w,{key:b.action,xs:12,sm:12,md:12,lg:6,xl:6,xxl:6,class:"list-col"},{default:t(()=>[ie("div",Je,[l(G,{title:b.title,bordered:!1,"header-style":{padding:"16px"},"body-style":{padding:"0px 16px"}},{extra:t(()=>[b.open!==void 0?(u(),n(O,{key:0,modelValue:b.open,"onUpdate:modelValue":h=>b.open=h,onChange:h=>W(b)},null,8,["modelValue","onUpdate:modelValue","onChange"])):_("",!0)]),actions:t(()=>[b.reset?(u(),n(H,{key:0,onClick:h=>x(b)},{default:t(()=>[A(U(s.$t("button.reset")),1)]),_:2},1032,["onClick"])):_("",!0),b.config?(u(),n(H,{key:1,type:"primary",onClick:h=>R(b)},{default:t(()=>[A(U(s.$t("button.config")),1)]),_:2},1032,["onClick"])):_("",!0)]),skeleton:t(()=>[l(g,{animation:!0},{default:t(()=>[l(r,{widths:["50%","50%","100%","40%"],rows:4}),l(r,{widths:["40%"],rows:1})]),_:1})]),default:t(()=>[l(I,null,{description:t(()=>[A(U(b.description)+" ",1),l(T,{style:{"margin-top":"10px"},data:b.data,layout:"inline-horizontal",column:2},null,8,["data"])]),_:2},1024)]),_:2},1032,["title"])])]),_:2},1024))),128))]),_:1}),l(Be,{visible:m.value,"onUpdate:visible":D[1]||(D[1]=b=>m.value=b),title:s.$t(q.value),onCancel:j,onBeforeOk:P},{default:t(()=>[l(S,{ref_key:"configForm",ref:C,model:e.value,"auto-label-width":"",style:{"max-height":"300px"}},{default:t(()=>[(u(!0),V(z,null,M(e.value.user_shield_error.errors,(b,h)=>Y((u(),n(a,{key:h,field:`user_shield_error.errors[${h}]`,label:`${h+1}. `+s.$t("sys.config.label.user_shield_error.errors"),rules:[{required:!0,message:s.$t("sys.config.error.user_shield_error.errors.required")}],"label-col-style":{padding:"0 16px 2px 0"}},{default:t(()=>[l($,{modelValue:e.value.user_shield_error.errors[h],"onUpdate:modelValue":K=>e.value.user_shield_error.errors[h]=K,placeholder:s.$t("sys.config.placeholder.user_shield_error.errors"),"allow-clear":"",style:{width:"75%","margin-right":"5px"}},null,8,["modelValue","onUpdate:modelValue","placeholder"]),l(H,{type:"primary",shape:"circle",style:{margin:"0 10px 0 10px"},onClick:D[0]||(D[0]=K=>o())},{default:t(()=>[l(J)]),_:1}),l(H,{type:"secondary",shape:"circle",onClick:K=>i(h)},{default:t(()=>[l(X)]),_:2},1032,["onClick"])]),_:2},1032,["field","label","rules"])),[[Z,e.value.action==="user_shield_error"]])),128))]),_:1},8,["model"])]),_:1},8,["visible","title"])])}}});const Ve=ee(je,[["__scopeId","data-v-3ba125e7"]]),Qe={class:"list-wrap",style:{"margin-top":"10px"}},We={class:"card-wrap"},Xe={name:"Task"},Ye=le({...Xe,setup(L){const{proxy:p}=he(),{setLoading:d}=Fe(!0),{t:y}=ne(),m=v(!1),q=v(""),C=v(),e=v({statistics:{}}),R=async o=>{q.value=y(`sys.config.item.title.${o.action}`),e.value.action=o.action,m.value=!0},x=async o=>{p.$modal.warning({title:"\u8B66\u544A",titleAlign:"center",content:`\u662F\u5426\u786E\u5B9A\u91CD\u7F6E${y(`sys.config.item.title.${o.action}`)}?`,hideCancel:!1,onOk:()=>{Q(o)}})},P=async o=>{var s;if(await((s=C.value)==null?void 0:s.validate())){m.value=!0,o(!1);return}d(!0);try{await Ce(e.value),o(),p.$message.success("\u64CD\u4F5C\u6210\u529F"),F()}catch{o(!1)}finally{d(!1)}},j=()=>{m.value=!1},Q=async o=>{d(!0);try{await ke({action:o.action}),p.$message.success("\u64CD\u4F5C\u6210\u529F"),F()}finally{d(!1)}},W=async o=>{d(!0);try{await Ee({action:o.action,open:o.open||!1}),p.$message.success("\u64CD\u4F5C\u6210\u529F"),F()}finally{d(!1)}},k=v({}),N=v({}),F=async()=>{const{data:o}=await $e();k.value=o,e.value.statistics=o.statistics,N.value=[{action:"statistics",title:y("sys.config.item.title.statistics"),description:"\u7EDF\u8BA1\u4EEA\u8868\u76D8\u4E0A\u7684\u5404\u7C7B\u6570\u636E",open:k.value.statistics.open,config:!0,reset:!0}]};return F(),(o,i)=>{const s=ce,D=de,O=_e,T=pe,I=me,H=fe,r=te,g=oe,G=re,w=ye,E=ge,$=De,J=ve,X=be;return u(),V("div",Qe,[l(G,{class:"list-row",gutter:24},{default:t(()=>[(u(!0),V(z,null,M(N.value,a=>(u(),n(g,{key:a.action,xs:12,sm:12,md:12,lg:6,xl:6,xxl:6,class:"list-col"},{default:t(()=>[ie("div",We,[l(r,{title:a.title,bordered:!1,"header-style":{padding:"16px"},"body-style":{padding:"0px 16px"}},{extra:t(()=>[a.open!==void 0?(u(),n(s,{key:0,modelValue:a.open,"onUpdate:modelValue":S=>a.open=S,onChange:S=>W(a)},null,8,["modelValue","onUpdate:modelValue","onChange"])):_("",!0)]),actions:t(()=>[a.reset?(u(),n(T,{key:0,onClick:S=>x(a)},{default:t(()=>[A(U(o.$t("button.reset")),1)]),_:2},1032,["onClick"])):_("",!0),a.config?(u(),n(T,{key:1,type:"primary",onClick:S=>R(a)},{default:t(()=>[A(U(o.$t("button.config")),1)]),_:2},1032,["onClick"])):_("",!0)]),skeleton:t(()=>[l(H,{animation:!0},{default:t(()=>[l(I,{widths:["50%","50%","100%","40%"],rows:4}),l(I,{widths:["40%"],rows:1})]),_:1})]),default:t(()=>[l(O,null,{description:t(()=>[A(U(a.description)+" ",1),l(D,{style:{"margin-top":"10px"},data:a.data,layout:"inline-horizontal",column:2},null,8,["data"])]),_:2},1024)]),_:2},1032,["title"])])]),_:2},1024))),128))]),_:1}),l(X,{visible:m.value,"onUpdate:visible":i[3]||(i[3]=a=>m.value=a),title:o.$t(q.value),onCancel:j,onBeforeOk:P},{default:t(()=>[l(J,{ref_key:"configForm",ref:C,model:e.value,"auto-label-width":""},{default:t(()=>[e.value.action==="statistics"?(u(),n(E,{key:0,field:"statistics.cron",label:o.$t("sys.config.label.statistics.cron"),rules:[{required:!0,message:o.$t("sys.config.error.statistics.cron.required")}]},{default:t(()=>[l(w,{modelValue:e.value.statistics.cron,"onUpdate:modelValue":i[0]||(i[0]=a=>e.value.statistics.cron=a),placeholder:o.$t("sys.config.placeholder.statistics.cron"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"])):_("",!0),e.value.action==="statistics"?(u(),n(E,{key:1,field:"statistics.limit",label:o.$t("sys.config.label.statistics.limit"),rules:[{required:!0,message:o.$t("sys.config.error.statistics.limit.required")}]},{default:t(()=>[l($,{modelValue:e.value.statistics.limit,"onUpdate:modelValue":i[1]||(i[1]=a=>e.value.statistics.limit=a),placeholder:o.$t("sys.config.placeholder.statistics.limit"),min:10,"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"])):_("",!0),e.value.action==="statistics"?(u(),n(E,{key:2,field:"statistics.lock_minutes",label:o.$t("sys.config.label.statistics.lock_minutes"),rules:[{required:!0,message:o.$t("sys.config.error.statistics.lock_minutes.required")}]},{default:t(()=>[l($,{modelValue:e.value.statistics.lock_minutes,"onUpdate:modelValue":i[2]||(i[2]=a=>e.value.statistics.lock_minutes=a),placeholder:o.$t("sys.config.placeholder.statistics.lock_minutes"),min:1,"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"])):_("",!0)]),_:1},8,["model"])]),_:1},8,["visible","title"])])}}});const Ae=ee(Ye,[["__scopeId","data-v-dc19feca"]]),Ze={class:"list-wrap",style:{"margin-top":"10px"}},el={class:"card-wrap"},ll={name:"API"},tl=le({...ll,setup(L){const{proxy:p}=he(),{setLoading:d}=Fe(!0),{t:y}=ne(),m=v(!1),q=v(""),C=v(),e=v({base:{},log:{},auto_disabled_error:{},not_retry_error:{},not_shield_error:{}}),R=async r=>{r.action==="auto_disabled_error"&&e.value.auto_disabled_error.errors.length===0&&s(),r.action==="not_retry_error"&&e.value.not_retry_error.errors.length===0&&O(),r.action==="not_shield_error"&&e.value.not_shield_error.errors.length===0&&I(),q.value=y(`sys.config.item.title.${r.action}`),e.value.action=r.action,m.value=!0},x=async r=>{p.$modal.warning({title:"\u8B66\u544A",titleAlign:"center",content:`\u662F\u5426\u786E\u5B9A\u91CD\u7F6E${y(`sys.config.item.title.${r.action}`)}?`,hideCancel:!1,onOk:()=>{Q(r)}})},P=async r=>{var G;if(await((G=C.value)==null?void 0:G.validate())){m.value=!0,r(!1);return}d(!0);try{await Ce(e.value),r(),p.$message.success("\u64CD\u4F5C\u6210\u529F"),F()}catch{r(!1)}finally{d(!1)}},j=()=>{m.value=!1,e.value.auto_disabled_error.errors.length>0&&!e.value.auto_disabled_error.errors[e.value.auto_disabled_error.errors.length-1]&&D(e.value.auto_disabled_error.errors.length-1),e.value.not_retry_error.errors.length>0&&!e.value.not_retry_error.errors[e.value.not_retry_error.errors.length-1]&&T(e.value.not_retry_error.errors.length-1),e.value.not_shield_error.errors.length>0&&!e.value.not_shield_error.errors[e.value.not_shield_error.errors.length-1]&&H(e.value.not_shield_error.errors.length-1)},Q=async r=>{d(!0);try{await ke({action:r.action}),p.$message.success("\u64CD\u4F5C\u6210\u529F"),F()}finally{d(!1)}},W=async r=>{d(!0);try{await Ee({action:r.action,open:r.open||!1}),p.$message.success("\u64CD\u4F5C\u6210\u529F"),F()}finally{d(!1)}},k=v({}),N=v({}),F=async()=>{const{data:r}=await $e();k.value=r,e.value.base=r.base,e.value.log=r.log,e.value.auto_disabled_error=r.auto_disabled_error,e.value.not_retry_error=r.not_retry_error,e.value.not_shield_error=r.not_shield_error,N.value=[{action:"base",title:y("sys.config.item.title.base"),description:"\u914D\u7F6E\u9519\u8BEF\u91CD\u8BD5\u6B21\u6570\u548C\u5404\u7C7B\u9519\u8BEF\u7981\u7528\u6B21\u6570\u7B49, \u9519\u8BEF\u91CD\u8BD5\u6B21\u6570N > 0 \u91CD\u8BD5 N \u6B21, N < 0 \u91CD\u8BD5\u6240\u6709key\u4E00\u8F6E, N = 0 \u4E0D\u91CD\u8BD5, \u9519\u8BEF\u6B21\u6570\u6BCF\u59290\u70B9\u4F1A\u81EA\u52A8\u91CD\u7F6E, \u6CE8: \u4EE3\u7406\u5BC6\u94A5\u9519\u8BEF\u65F6, \u4E5F\u4F1A\u8BB0\u5F55\u6A21\u578B\u4EE3\u7406\u9519\u8BEF\u6B21\u6570",config:!0,reset:!0},{action:"log",title:y("sys.config.item.title.log"),description:"\u8C03\u7528\u65E5\u5FD7\u8BB0\u5F55\u5185\u5BB9, prompt: \u63D0\u95EE, completion: \u56DE\u7B54, messages: \u4E0A\u4E0B\u6587, image: \u591A\u6A21\u6001\u8BC6\u56FE\u7684BASE64\u56FE\u50CF\u6570\u636E",open:k.value.log.open,config:!0,reset:!0},{action:"auto_disabled_error",title:y("sys.config.item.title.auto_disabled_error"),description:"\u8C03\u7528\u62A5\u9519\u65F6, \u5305\u542B\u6709\u914D\u7F6E\u9519\u8BEF\u5185\u5BB9\u65F6\u5219\u81EA\u52A8\u4F1A\u7981\u7528\u5BC6\u94A5\u6216\u6A21\u578B\u4EE3\u7406\u7B49, \u4E3A\u7A7A\u5219\u4E0D\u4F1A\u81EA\u52A8\u7981\u7528(\u8FBE\u5230\u9519\u8BEF\u6B21\u6570\u4E0A\u9650\u9664\u5916)",open:k.value.auto_disabled_error.open,config:!0,reset:!0},{action:"not_retry_error",title:y("sys.config.item.title.not_retry_error"),description:"\u8C03\u7528\u62A5\u9519\u65F6, \u5305\u542B\u6709\u914D\u7F6E\u9519\u8BEF\u5185\u5BB9\u65F6\u5219\u4E0D\u4F1A\u81EA\u52A8\u91CD\u8BD5, \u4E3A\u7A7A\u5219\u4F1A\u81EA\u52A8\u91CD\u8BD5",open:k.value.not_retry_error.open,config:!0,reset:!0},{action:"not_shield_error",title:y("sys.config.item.title.not_shield_error"),description:"\u8C03\u7528\u62A5\u9519\u65F6, \u5305\u542B\u6709\u914D\u7F6E\u9519\u8BEF\u5185\u5BB9\u65F6\u5219\u4F1A\u5C06\u9519\u8BEF\u5185\u5BB9\u8FD4\u56DE\u7ED9\u8C03\u7528\u65B9, \u4E3A\u7A7A\u5219\u5C4F\u853D\u6240\u6709\u9519\u8BEF",open:k.value.not_shield_error.open,config:!0,reset:!0},{action:"reset_api_error",title:y("sys.config.item.title.reset_api_error"),description:"\u5F53\u9519\u8BEF\u6B21\u6570\u8FBE\u5230\u914D\u7F6E\u4E0A\u9650\u65F6\u53EF\u624B\u52A8\u8FDB\u884C\u91CD\u7F6E, \u91CD\u7F6E\u8FC7\u7A0B\u53EF\u80FD\u4F1A\u9020\u6210\u7CFB\u7EDF\u7684\u77ED\u6682\u4E0D\u53EF\u7528(\u4E00\u822C\u51E0\u79D2\u949F), \u8BF7\u8C28\u614E\u64CD\u4F5C",reset:!0}]};F();const o=()=>{e.value.log.records.push("")},i=r=>{e.value.log.records.length>1&&e.value.log.records.splice(r,1)},s=()=>{e.value.auto_disabled_error.errors.push("")},D=r=>{e.value.auto_disabled_error.errors.splice(r,1)},O=()=>{e.value.not_retry_error.errors.push("")},T=r=>{e.value.not_retry_error.errors.splice(r,1)},I=()=>{e.value.not_shield_error.errors.push("")},H=r=>{e.value.not_shield_error.errors.splice(r,1)};return(r,g)=>{const G=ce,w=de,E=_e,$=pe,J=me,X=fe,a=te,S=oe,Be=re,b=De,h=ge,K=ye,ae=qe,ue=Se,Ne=ve,Oe=be;return u(),V("div",Ze,[l(Be,{class:"list-row",gutter:24},{default:t(()=>[(u(!0),V(z,null,M(N.value,f=>(u(),n(S,{key:f.action,xs:12,sm:12,md:12,lg:6,xl:6,xxl:6,class:"list-col"},{default:t(()=>[ie("div",el,[l(a,{title:f.title,bordered:!1,"header-style":{padding:"16px"},"body-style":{padding:"0px 16px"}},{extra:t(()=>[f.open!==void 0?(u(),n(G,{key:0,modelValue:f.open,"onUpdate:modelValue":c=>f.open=c,onChange:c=>W(f)},null,8,["modelValue","onUpdate:modelValue","onChange"])):_("",!0)]),actions:t(()=>[f.reset?(u(),n($,{key:0,onClick:c=>x(f)},{default:t(()=>[A(U(r.$t("button.reset")),1)]),_:2},1032,["onClick"])):_("",!0),f.config?(u(),n($,{key:1,type:"primary",onClick:c=>R(f)},{default:t(()=>[A(U(r.$t("button.config")),1)]),_:2},1032,["onClick"])):_("",!0)]),skeleton:t(()=>[l(X,{animation:!0},{default:t(()=>[l(J,{widths:["50%","50%","100%","40%"],rows:4}),l(J,{widths:["40%"],rows:1})]),_:1})]),default:t(()=>[l(E,null,{description:t(()=>[A(U(f.description)+" ",1),l(w,{style:{"margin-top":"10px"},data:f.data,layout:"inline-horizontal",column:2},null,8,["data"])]),_:2},1024)]),_:2},1032,["title"])])]),_:2},1024))),128))]),_:1}),l(Oe,{visible:m.value,"onUpdate:visible":g[8]||(g[8]=f=>m.value=f),title:r.$t(q.value),onCancel:j,onBeforeOk:P},{default:t(()=>[l(Ne,{ref_key:"configForm",ref:C,model:e.value,"auto-label-width":"",style:{"max-height":"300px"}},{default:t(()=>[e.value.action==="base"?(u(),n(h,{key:0,field:"base.err_retry",label:r.$t("sys.config.label.base.err_retry"),rules:[{required:!0,message:r.$t("sys.config.error.base.err_retry.required")}]},{default:t(()=>[l(b,{modelValue:e.value.base.err_retry,"onUpdate:modelValue":g[0]||(g[0]=f=>e.value.base.err_retry=f),placeholder:r.$t("sys.config.placeholder.base.err_retry"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"])):_("",!0),e.value.action==="base"?(u(),n(h,{key:1,field:"base.model_key_err_disable",label:r.$t("sys.config.label.base.model_key_err_disable"),rules:[{required:!0,message:r.$t("sys.config.error.base.model_key_err_disable.required")}]},{default:t(()=>[l(b,{modelValue:e.value.base.model_key_err_disable,"onUpdate:modelValue":g[1]||(g[1]=f=>e.value.base.model_key_err_disable=f),placeholder:r.$t("sys.config.placeholder.base.model_key_err_disable"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"])):_("",!0),e.value.action==="base"?(u(),n(h,{key:2,field:"base.model_agent_err_disable",label:r.$t("sys.config.label.base.model_agent_err_disable"),rules:[{required:!0,message:r.$t("sys.config.error.base.model_agent_err_disable.required")}]},{default:t(()=>[l(b,{modelValue:e.value.base.model_agent_err_disable,"onUpdate:modelValue":g[2]||(g[2]=f=>e.value.base.model_agent_err_disable=f),placeholder:r.$t("sys.config.placeholder.base.model_agent_err_disable"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"])):_("",!0),e.value.action==="base"?(u(),n(h,{key:3,field:"base.model_agent_key_err_disable",label:r.$t("sys.config.label.base.model_agent_key_err_disable"),rules:[{required:!0,message:r.$t("sys.config.error.base.model_agent_key_err_disable.required")}]},{default:t(()=>[l(b,{modelValue:e.value.base.model_agent_key_err_disable,"onUpdate:modelValue":g[3]||(g[3]=f=>e.value.base.model_agent_key_err_disable=f),placeholder:r.$t("sys.config.placeholder.base.model_agent_key_err_disable"),"allow-clear":""},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"])):_("",!0),(u(!0),V(z,null,M(e.value.log.records,(f,c)=>Y((u(),n(h,{key:c,field:`log.records[${c}]`,label:`${c+1}. `+r.$t("sys.config.label.log.records"),rules:[{required:!0,message:r.$t("sys.config.error.log.records.required")}],"label-col-style":{padding:"0 16px 2px 0"}},{default:t(()=>[l(K,{modelValue:e.value.log.records[c],"onUpdate:modelValue":B=>e.value.log.records[c]=B,placeholder:r.$t("sys.config.placeholder.log.records"),"allow-clear":"",style:{width:"75%","margin-right":"5px"}},null,8,["modelValue","onUpdate:modelValue","placeholder"]),l($,{type:"primary",shape:"circle",style:{margin:"0 10px 0 10px"},onClick:g[4]||(g[4]=B=>o())},{default:t(()=>[l(ae)]),_:1}),l($,{type:"secondary",shape:"circle",onClick:B=>i(c)},{default:t(()=>[l(ue)]),_:2},1032,["onClick"])]),_:2},1032,["field","label","rules"])),[[Z,e.value.action==="log"]])),128)),(u(!0),V(z,null,M(e.value.auto_disabled_error.errors,(f,c)=>Y((u(),n(h,{key:c,field:`auto_disabled_error.errors[${c}]`,label:`${c+1}. `+r.$t("sys.config.label.auto_disabled_error.errors"),rules:[{required:!0,message:r.$t("sys.config.error.auto_disabled_error.errors.required")}],"label-col-style":{padding:"0 16px 2px 0"}},{default:t(()=>[l(K,{modelValue:e.value.auto_disabled_error.errors[c],"onUpdate:modelValue":B=>e.value.auto_disabled_error.errors[c]=B,placeholder:r.$t("sys.config.placeholder.auto_disabled_error.errors"),"allow-clear":"",style:{width:"75%","margin-right":"5px"}},null,8,["modelValue","onUpdate:modelValue","placeholder"]),l($,{type:"primary",shape:"circle",style:{margin:"0 10px 0 10px"},onClick:g[5]||(g[5]=B=>s())},{default:t(()=>[l(ae)]),_:1}),l($,{type:"secondary",shape:"circle",onClick:B=>D(c)},{default:t(()=>[l(ue)]),_:2},1032,["onClick"])]),_:2},1032,["field","label","rules"])),[[Z,e.value.action==="auto_disabled_error"]])),128)),(u(!0),V(z,null,M(e.value.not_retry_error.errors,(f,c)=>Y((u(),n(h,{key:c,field:`not_retry_error.errors[${c}]`,label:`${c+1}. `+r.$t("sys.config.label.not_retry_error.errors"),rules:[{required:!0,message:r.$t("sys.config.error.not_retry_error.errors.required")}],"label-col-style":{padding:"0 16px 2px 0"}},{default:t(()=>[l(K,{modelValue:e.value.not_retry_error.errors[c],"onUpdate:modelValue":B=>e.value.not_retry_error.errors[c]=B,placeholder:r.$t("sys.config.placeholder.not_retry_error.errors"),"allow-clear":"",style:{width:"75%","margin-right":"5px"}},null,8,["modelValue","onUpdate:modelValue","placeholder"]),l($,{type:"primary",shape:"circle",style:{margin:"0 10px 0 10px"},onClick:g[6]||(g[6]=B=>O())},{default:t(()=>[l(ae)]),_:1}),l($,{type:"secondary",shape:"circle",onClick:B=>T(c)},{default:t(()=>[l(ue)]),_:2},1032,["onClick"])]),_:2},1032,["field","label","rules"])),[[Z,e.value.action==="not_retry_error"]])),128)),(u(!0),V(z,null,M(e.value.not_shield_error.errors,(f,c)=>Y((u(),n(h,{key:c,field:`not_shield_error.errors[${c}]`,label:`${c+1}. `+r.$t("sys.config.label.not_shield_error.errors"),rules:[{required:!0,message:r.$t("sys.config.error.not_shield_error.errors.required")}],"label-col-style":{padding:"0 16px 2px 0"}},{default:t(()=>[l(K,{modelValue:e.value.not_shield_error.errors[c],"onUpdate:modelValue":B=>e.value.not_shield_error.errors[c]=B,placeholder:r.$t("sys.config.placeholder.not_shield_error.errors"),"allow-clear":"",style:{width:"75%","margin-right":"5px"}},null,8,["modelValue","onUpdate:modelValue","placeholder"]),l($,{type:"primary",shape:"circle",style:{margin:"0 10px 0 10px"},onClick:g[7]||(g[7]=B=>I())},{default:t(()=>[l(ae)]),_:1}),l($,{type:"secondary",shape:"circle",onClick:B=>H(c)},{default:t(()=>[l(ue)]),_:2},1032,["onClick"])]),_:2},1032,["field","label","rules"])),[[Z,e.value.action==="not_shield_error"]])),128))]),_:1},8,["model"])]),_:1},8,["visible","title"])])}}});const Ue=ee(tl,[["__scopeId","data-v-deac8740"]]),ol={class:"container"},rl={name:"SysConfig"},al=le({...rl,setup(L){return(p,d)=>{const y=Te,m=He,q=Le,C=Re,e=xe,R=oe,x=re,P=te;return u(),V("div",ol,[l(q,{class:"container-breadcrumb"},{default:t(()=>[l(m,null,{default:t(()=>[l(y)]),_:1}),l(m,null,{default:t(()=>[A(U(p.$t("menu.sys")),1)]),_:1}),l(m,null,{default:t(()=>[A(U(p.$t("menu.sys.config")),1)]),_:1})]),_:1}),l(x,{gutter:20,align:"stretch"},{default:t(()=>[l(R,{span:24},{default:t(()=>[l(P,{class:"general-card",bordered:!1,"header-style":{padding:"20px"},"body-style":{padding:"25px 20px 20px 14px"}},{default:t(()=>[l(x,{justify:"space-between"},{default:t(()=>[l(R,{span:24},{default:t(()=>[l(e,{"default-active-tab":0,type:"rounded"},{default:t(()=>[l(C,{key:"0",title:p.$t("sys.config.tab.title.all")},{default:t(()=>[l(we),l(Ve),l(Ae),l(Ue)]),_:1},8,["title"]),l(C,{key:"1",title:p.$t("sys.config.tab.title.global")},{default:t(()=>[l(we)]),_:1},8,["title"]),l(C,{key:"2",title:p.$t("sys.config.tab.title.general")},{default:t(()=>[l(Ve)]),_:1},8,["title"]),l(C,{key:"3",title:p.$t("sys.config.tab.title.task")},{default:t(()=>[l(Ae)]),_:1},8,["title"]),l(C,{key:"4",title:p.$t("sys.config.tab.title.api")},{default:t(()=>[l(Ue)]),_:1},8,["title"])]),_:1})]),_:1})]),_:1})]),_:1})]),_:1})]),_:1})])}}});const vl=ee(al,[["__scopeId","data-v-2861651e"]]);export{vl as default};
