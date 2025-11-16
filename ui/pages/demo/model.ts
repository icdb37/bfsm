export interface ListSelecter{
	getTitle:()=>string,
	getNote:()=>string
}

class Goods implements ListSelecter{
	name:string;
	spec:string;
	size:string;
	
	constructor(name:string,spec:string,size:string) {
		this.name = name;
		this.spec = spec;
		this.size = size;
	}
	
	getTitle():string{
		return this.name;
	}
	getNote():string{
		return this.spec + "-" + this.size;
	}
}

export const examples:ListSelecter[] = [
	new Goods("a1", "a2", "a3"),
	new Goods("b1", "b2", "a3"),
	new Goods("c1", "b2", "a3"),
	new Goods("d1", "b2", "a3"),
	new Goods("e1", "b2", "a3"),
	new Goods("f1", "b2", "a3"),
	new Goods("g1", "b2", "a3"),
	new Goods("h1", "b2", "a3"),
	new Goods("i1", "b2", "a3"),
	new Goods("j1", "b2", "a3"),
]