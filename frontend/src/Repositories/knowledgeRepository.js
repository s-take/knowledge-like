import Repository from './Repository'

const resource = '/knowledge'

export default {
    get() {
        return Repository.get(`${resource}`);
    },
    createKnowledge(payload) {
        return Repository.post(`${resource}`, payload);
    }
};