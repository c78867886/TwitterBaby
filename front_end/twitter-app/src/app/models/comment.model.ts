export interface Comment {
    id ?:string;
    content: string;
    timestamp: string;
    commentnum ?: number;
    commentowner : string;
}