import axios from "axios";
import type { SearchRequestType } from "../types/Email.type";

 export async function listIndex() {
      const response = await axios.get('http://localhost:3000/list'); 
      return response.data; 
  }

  export async function deleteIndex(nameIndex: string){
    const response = await axios.delete('http://localhost:3000/emails', {params:{
      index_name: nameIndex
    }})

    console.log(response.data);
    
    return response.data
  }
  
  export async function searchEmail(params:SearchRequestType){
    const response = await axios.post('http://localhost:3000/search',{
        query: params.query || "",
        limit: params.limit || 50,
        offset: params.offset || 0,
        start_date: params.startDate || "",
        end_date: params.endDate || "",
        name_index: params.nameIndex,
      
    })
    return response.data

  }
  