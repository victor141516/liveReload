# liveReload
Golang liveReload for frontend development. Usage: ./liveReload [path] [port]

## Frontend code
```
// liveReload poller begin
const poller = setInterval(() => {
    fetch('http://127.0.0.1:15555', {mode: 'cors'}).then(res => res.text().then(text => {
        if (text === '1') {
            location.reload();
        }
    })).catch(() => {
        console.log('liveReload not detected, poller stopped.');clearInterval(poller); // eslint-disable-line
    });
}, 1000);
// liveReload poller end
```
