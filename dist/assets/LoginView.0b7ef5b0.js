import{c as p,o as u,a as m,w as k,v as S,u as i,b as c,i as v,t as f,n as b,d as w,e as g,r as x,f as $,g as B,h as n,j as s,_ as C,k as F,l as N,m as y,p as _,q as P,s as U,x as L}from"./index.e2e5f540.js";import{_ as R,a as j}from"./LayoutGuest.060751be.js";const q=["type","name","value"],A=c("span",{class:"check"},null,-1),D={class:"pl-2"},T={__name:"FormCheckRadio",props:{name:{type:String,required:!0},type:{type:String,default:"checkbox",validator:e=>["checkbox","radio","switch"].includes(e)},label:{type:String,default:null},modelValue:{type:[Array,String,Number,Boolean],default:null},inputValue:{type:[String,Number,Boolean],required:!0}},emits:["update:modelValue"],setup(e,{emit:t}){const r=e,l=p({get:()=>r.modelValue,set:d=>{t("update:modelValue",d)}}),a=p(()=>r.type==="radio"?"radio":"checkbox");return(d,o)=>(u(),m("label",{class:b(e.type)},[k(c("input",{"onUpdate:modelValue":o[0]||(o[0]=V=>v(l)?l.value=V:null),type:i(a),name:e.name,value:e.inputValue},null,8,q),[[S,i(l)]]),A,c("span",D,f(e.label),1)],2))}},M={class:"mb-6 last:mb-0"},O=["for"],z={key:1,class:"text-xs text-gray-500 dark:text-slate-400 mt-1"},h={__name:"FormField",props:{label:{type:String,default:null},labelFor:{type:String,default:null},help:{type:String,default:null}},setup(e){const t=w(),r=p(()=>{const l=[],a=t.default().length;return a>1&&l.push("grid grid-cols-1 gap-3"),a===2&&l.push("md:grid-cols-2"),l});return(l,a)=>(u(),m("div",M,[e.label?(u(),m("label",{key:0,for:e.labelFor,class:"block font-bold mb-2"},f(e.label),9,O)):g("",!0),c("div",{class:b(i(r))},[x(l.$slots,"default")],2),e.help?(u(),m("div",z,f(e.help),1)):g("",!0)]))}},G={__name:"LoginView",setup(e){const t=$({remember:!0});P();const r=()=>{fetch("/api/login",{method:"POST",body:JSON.stringify({name:t.login,password:t.pass}),headers:{"Content-type":"application/json; charset=UTF-8"}}).then(l=>l.json()).then(l=>console.log(l))};return(l,a)=>(u(),B(R,null,{default:n(()=>[s(j,{bg:"purplePink"},{default:n(({cardClass:d})=>[s(C,{class:b(d),"is-form":"",onSubmit:F(r,["prevent"])},{footer:n(()=>[s(N,null,{default:n(()=>[s(y,{type:"submit",color:"info",label:"Login"}),s(y,{to:"/",color:"info",outline:"",label:"Back"})]),_:1})]),default:n(()=>[s(h,{label:"Login",help:"Please enter your login"},{default:n(()=>[s(_,{modelValue:t.login,"onUpdate:modelValue":a[0]||(a[0]=o=>t.login=o),icon:i(U),name:"login",autocomplete:"username"},null,8,["modelValue","icon"])]),_:1}),s(h,{label:"Password",help:"Please enter your password"},{default:n(()=>[s(_,{modelValue:t.pass,"onUpdate:modelValue":a[1]||(a[1]=o=>t.pass=o),icon:i(L),type:"password",name:"password",autocomplete:"current-password"},null,8,["modelValue","icon"])]),_:1}),s(T,{modelValue:t.remember,"onUpdate:modelValue":a[2]||(a[2]=o=>t.remember=o),name:"remember",label:"Remember","input-value":!0},null,8,["modelValue"])]),_:2},1032,["class","onSubmit"])]),_:1})]),_:1}))}};export{G as default};