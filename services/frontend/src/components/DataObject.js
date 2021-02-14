import { observer, inject } from 'mobx-react';


const DataObject = inject('objectStore')(observer(({objectStore}) => {
    const objectList = objectStore.objects.map(obj => (
        <div className="object" key={obj.num}>
            <div className="object-num">{obj.num}</div>
            <div className="object-id">{obj.id}</div>
            <div className="object-text">{obj.text}</div>
        </div>
    ))
    return (
        <div className='object-list'>
            { objectList }
        </div>
    )
}));

export default DataObject