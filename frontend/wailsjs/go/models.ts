export namespace main {
	
	export class GitHubUser {
	    login: string;
	    id: number;
	    avatar_url: string;
	    html_url: string;
	    name: string;
	    company: string;
	    blog: string;
	    location: string;
	    bio: string;
	    public_repos: number;
	    followers: number;
	    following: number;
	    created_at: string;
	
	    static createFrom(source: any = {}) {
	        return new GitHubUser(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.login = source["login"];
	        this.id = source["id"];
	        this.avatar_url = source["avatar_url"];
	        this.html_url = source["html_url"];
	        this.name = source["name"];
	        this.company = source["company"];
	        this.blog = source["blog"];
	        this.location = source["location"];
	        this.bio = source["bio"];
	        this.public_repos = source["public_repos"];
	        this.followers = source["followers"];
	        this.following = source["following"];
	        this.created_at = source["created_at"];
	    }
	}
	export class Repository {
	    name: string;
	    stargazers_count: number;
	    language: string;
	    created_at: string;
	
	    static createFrom(source: any = {}) {
	        return new Repository(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.stargazers_count = source["stargazers_count"];
	        this.language = source["language"];
	        this.created_at = source["created_at"];
	    }
	}

}

