import{c as i,o as t,a as n,r as o,n as c,u as l,y as d,z as u,A as p,B as _,b as k}from"./index.f9d5912f.js";const y={__name:"SectionFullScreen",props:{bg:{type:String,required:!0,validator:e=>["purplePink","pinkRed"].includes(e)}},setup(e){const s=e,r=i(()=>{if(d().darkMode)return u;switch(s.bg){case"purplePink":return _;case"pinkRed":return p}return""});return(a,m)=>(t(),n("div",{class:c(["flex min-h-screen items-center justify-center",l(r)])},[o(a.$slots,"default",{cardClass:"w-11/12 md:w-7/12 lg:w-6/12 xl:w-4/12 shadow-2xl"})],2))}},g={class:"bg-gray-50 dark:bg-slate-800 dark:text-slate-100"},B={__name:"LayoutGuest",setup(e){const s=d();return(r,a)=>(t(),n("div",{class:c({dark:l(s).darkMode})},[k("div",g,[o(r.$slots,"default")])],2))}};export{B as _,y as a};
