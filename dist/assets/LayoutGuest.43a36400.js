import{a as i,o as t,e as n,p as o,n as l,f as c,G as d,H as u,I as p,J as _,l as k}from"./index.0ced0e98.js";const S={__name:"SectionFullScreen",props:{bg:{type:String,required:!0,validator:e=>["purplePink","pinkRed"].includes(e)}},setup(e){const s=e,r=i(()=>{if(d().darkMode)return u;switch(s.bg){case"purplePink":return _;case"pinkRed":return p}return""});return(a,m)=>(t(),n("div",{class:l(["flex min-h-screen items-center justify-center",c(r)])},[o(a.$slots,"default",{cardClass:"w-11/12 md:w-7/12 lg:w-6/12 xl:w-4/12 shadow-2xl"})],2))}},g={class:"bg-gray-50 dark:bg-slate-800 dark:text-slate-100"},h={__name:"LayoutGuest",setup(e){const s=d();return(r,a)=>(t(),n("div",{class:l({dark:c(s).darkMode})},[k("div",g,[o(r.$slots,"default")])],2))}};export{h as _,S as a};