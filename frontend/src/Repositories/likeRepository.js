import Repository from './Repository'

const resource = '/like'

export default {
    get() {
        return Repository.get(`${resource}`);
    },
    getByUserID() {
        return Repository.get(`${resource}/my`);
    },
    getByKnowledgeID(id) {
        return Repository.get(`${resource}/${id}`);
    },
    createLike(payload) {
        return Repository.post(`${resource}`, payload);
    },
    deleteLike() {
        return Repository.delete(`${resource}/my`);
    }
};