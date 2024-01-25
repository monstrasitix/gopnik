type Keys = {
    version: string;
    schema: {
        foo: string;
        bar: number;
    };
};

class Store<T extends {}> {
    constructor(private store: Storage) {}

    find<T1 extends keyof T>(name: T1): string | null {
        return this.store.getItem(name.toString());
    }
}

const store = new Store<Keys>(localStorage);

store.find("version");
