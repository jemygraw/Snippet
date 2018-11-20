const socket=new WebSocket('ws://localhost:9000');
socket.addEventListener('open',function(event){
    socket.send('hello server!');
});

socket.addEventListener('message',function(event){
    console.log('message received:',event.data);
});