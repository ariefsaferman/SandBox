export interface IBook {
    id: number; 
    title: string; 
    price: number; 
    description: string; 
    category: string; 
    image: string;
    rating:{
        rate: number; 
        count: number; 
    };
}

export interface IBookReq {
    title: string; 
    description: string;
    category: string; 
    image: string; 
    price: number; 
}

