export namespace wfmplatefficiency {
	
	export class Item {
	    name: string;
	    type: number;
	    standing: number;
	    WeightedAvgPrice: number;
	    AvgVol: number;
	
	    static createFrom(source: any = {}) {
	        return new Item(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.type = source["type"];
	        this.standing = source["standing"];
	        this.WeightedAvgPrice = source["WeightedAvgPrice"];
	        this.AvgVol = source["AvgVol"];
	    }
	}
	export class Vendor {
	    name: string;
	    items: Item[];
	
	    static createFrom(source: any = {}) {
	        return new Vendor(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.items = this.convertValues(source["items"], Item);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

