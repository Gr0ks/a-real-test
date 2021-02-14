import { action, makeObservable, observable } from 'mobx';
import throttle from "lodash/throttle";

const SERVER_ADR = 'http://localhost:8091'

class ObjectsStore {
    constructor() {
        this.objects = [];
        makeObservable(this, {
            objects: observable,
            __setObjects: action,
        });
        this.getObjects(50)
        window.addEventListener(
            'scroll', 
            throttle(() => {
                let windowRelativeBottom = document.documentElement.getBoundingClientRect().bottom;
                if (windowRelativeBottom < document.documentElement.clientHeight) {
                    window.objectStore.getObjects(20);
                }
            }, 1000),
        )
    }

    __setObjects(objectsArray) {
        const newObjects = this.objects.concat(objectsArray)
        this.objects = newObjects
    }

    async getObjects(count) {
        const resp = await fetch(
            `${SERVER_ADR}/api/getObjects?first=${this.objects.length}&count=${count}`
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