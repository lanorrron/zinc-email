import axios from "axios";
import type { SearchRequestType } from "../types/Email.type";
const BASE_URL = import.meta.env.VITE_BASE_URL

 export async function listIndex() {
      const response = await axios.get(`${BASE_URL}/list`); 
      return response.data; 
  }

  export async function deleteIndex(nameIndex: string){
    const response = await axios.delete(`${BASE_URL}/emails`, {params:{
      index_name: nameIndex
    }})

    console.log(response.data);
    
    return response.data
  }
  
  export async function searchEmail(params:SearchRequestType){
    const response = await axios.post(`${BASE_URL}/search`,{
        query: params.query?.trim() || "",
        limit: params.limit || 50,
        offset: params.offset || 0,
        start_date: params.startDate || "",
        end_date: params.endDate || "",
        name_index: params.nameIndex,
      
    })
    return response.data

  }
  