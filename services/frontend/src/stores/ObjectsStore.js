import { action, makeObservable, observable } from 'mobx'

const SERVER_ADR = 'http://localhost:8091'

class ObjectsStore {
    constructor() {
        this.objects = [];
        makeObservable(this, {
            objects: observable,
            __setObjects: action,
        });
    }

    __setObjects(objectsArray) {
        this.objects = objectsArray
    }

    async getObjects(firstObjectNumber, count) {
        const resp = await fetch(
            `${SERVER_ADR}/api/getObjects?first=${firstObjectNumber}&count=${count}`
        );
        if (resp.ok) {
            const respJson = await resp.json();
            console.log(respJson)
            this.__setObjects(respJson)
        }
    }
}

window.objectStore = new ObjectsStore();

export default window.objectStore