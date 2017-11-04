import { Injectable, EventEmitter } from '@angular/core';

@Injectable()
export class NotificationService {
  private socket: WebSocket;
  private listener: EventEmitter<any> = new EventEmitter();
  
  constructor() { 
    
  }

  public connect(id: string) {
    this.socket = new WebSocket("ws://localhost:1323/api/v1/ws/" + id);
    this.socket.onopen = event => {
      this.listener.emit({"type": "open", "data": event});
    }
    this.socket.onclose = event => {
      this.listener.emit({"type": "close", "data": event});
    }
    this.socket.onmessage = (event: any) => {
      this.listener.emit({"type": "message", "data": event.data});
    }
  }
 
  public readyState(): boolean {
    if (this.socket === undefined) {
      return false;
    }
    return this.socket.readyState === 1;
  }

  public send(data: string) {
    this.socket.send(data);
  }

  public close() {
    console.log("should close the connection!");
    this.socket.close();
  }

  public getEventListener() {
    return this.listener;
  }
  
}
