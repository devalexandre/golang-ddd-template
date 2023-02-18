mkdir -p internal/domain/user &&
touch internal/domain/user/repository.go  && 
touch internal/domain/user/service.go  &&   
touch internal/domain/user/repository_test.go  &&  
touch internal/domain/user/service_test.go  &&  
touch internal/domain/user/factory.go  &&  
touch internal/domain/user/factory_test.go  &&  
touch internal/domain/user/contracts.go  &&  
mkdir -p internal/domain/user/mocks &&
touch internal/domain/user/mocks/repository.go  &&  
touch internal/domain/user/mocks/service.go  &&  

mkdir -p internal/infra &&
mkdir -p internal/infra/mysql &&
touch internal/infra/mysql/connection.go  &&  
touch internal/infra/mysql/connection_test.go  &&  
touch internal/infra/mysql/contracts.go  &&  
mkdir -p internal/infra/mysql/mocks &&
touch internal/infra/mysql/mocks/connection.go  &&  

mkdir -p internal/helpers &&
touch internal/helpers/errors.go  &&  
touch internal/helpers/errors_test.go  &&  
mkdir -p internal/helpers/config &&
touch internal/helpers/config/config.go  