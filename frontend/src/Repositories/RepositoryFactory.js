import KnowledgeRepository from "./knowledgeRepository";
import LikeRepository from "./likeRepository";

const repositories = {
    knowledge: KnowledgeRepository,
    like: LikeRepository,
};

export const RepositoryFactory = {
    get: name => repositories[name]
}