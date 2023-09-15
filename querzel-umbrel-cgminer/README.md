# Gekko Science 2PAC & NEWPAC - Docker cgminer

Create `.env` file with key `POOL_USER` containing your mining address.
For details see [docker-compose docs](https://docs.docker.com/compose/environment-variables/#the-env-file) and
[ckpool docs](http://solo.ckpool.org/)

Start
```
docker-compose up --detach
```

Stop
```
docker-compose down
```

Rebuild
```
docker-compose build --force-rm --no-cache
```
