import{u as l,j as e,m as i,d as r,e as x,I as c,f as p,r as d,M as h,N as m,C as j,c as v,G as g,R as u,a as f,P as N}from"./index-70c75c17.js";const C=({movie:s})=>{const a=l();return e.exports.jsx(i.div,{initial:{scale:0,opacity:0},animate:{opacity:1,scale:1},transition:{type:"spring",stiffness:260,damping:40},children:e.exports.jsx(r,{style:{width:"100%",background:"#161616",color:"white",borderRadius:6,position:"relative"},className:" movie-card",children:e.exports.jsxs(r.Body,{children:[e.exports.jsx(x.exports.LazyLoadImage,{src:!s.poster_path||!s.backdrop_path?c:`${p}${s.backdrop_path}`,width:"100%",height:350,alt:"movie",effect:"blur",style:{objectFit:"cover"}}),e.exports.jsx(r.Title,{onClick:()=>a(`/tv/${s.id}`),className:"text-center mt-3",style:{cursor:"pointer"},children:s.name||s.title})]})})})};function b(){const{tvShows:s,tvGenres:a,handleTvGenres:o,tvTotalPages:n}=d.exports.useContext(h);return e.exports.jsxs(e.exports.Fragment,{children:[e.exports.jsx(m,{}),e.exports.jsxs(j,{className:"mt-4",children:[e.exports.jsx(i.div,{variants:v,initial:"hidden",animate:"visible",className:"genres d-flex flex-wrap ",style:{gap:"5px 15px"},children:a==null?void 0:a.map(t=>e.exports.jsx(g,{title:t.name,id:t.id,active:t.active,handleGenres:o},t.id))}),e.exports.jsx("div",{className:"wrapper mt-4",children:e.exports.jsx(u,{md:3,xs:1,lg:4,className:"g-4",children:s.map(t=>e.exports.jsx(f,{children:e.exports.jsx(C,{movie:t,tvShow:!0})},t.id))})}),e.exports.jsx("div",{className:"mt-5 d-flex justify-content-center",children:e.exports.jsx(N,{totalPages:n})})]})]})}export{b as default};
