// BaseURL 后端服务地址
export const BaseURL = process.env.VUE_APP_API_BASE_URL || "http://localhost:8080";


// SearchRequest 搜索请求
export class SearchRequest<T> {
  index: number;    // 分页索引
  size: number;     // 分页大小
  sorts: string[];  // 排序字段，如：-id：按id降序，id：按id升序
  query: T;         // 查询条件
  
  constructor(){
    this.index = 0;
    this.size = 10;
    this.sorts = [];
    this.query = {} as T;
  }
}

// SearchResponse 搜索应答
export class SearchResponse<T> {
  total: number;    // 总数
  data: T[];        // 列表
}


export interface Commodity {
  hash : string,     // 哈希
  name : string,     // 名称
  desc : string,     // 描述
  sepc : string,     // 规格
  size : string,     // 尺寸
  validity : number, // 有效期
  price : number,    // 价格
  count: number,     // 数量
}


export function FormatUTC(utc:string):string {
    let date = new Date(utc);
    let year = date.getFullYear();
    let month = String(date.getMonth() + 1).padStart(2, '0');
    let day = String(date.getDate()).padStart(2, '0');
    let hour = String(date.getHours()).padStart(2, '0');
    let minute = String(date.getMinutes()).padStart(2, '0');
    let second = String(date.getSeconds()).padStart(2, '0');
    return `${year}-${month}-${day} ${hour}:${minute}:${second}`;
}