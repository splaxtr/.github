const message = 'Hello from sample frontend';
console.log(message);
if (message.length === 0) {
  throw new Error('Message should not be empty');
}
