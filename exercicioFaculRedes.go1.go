package main

import "fmt"

struct no{

	float conteudo[2];
	struct no *esq;
	struct no *dir;
	
};

typedef struct no var;

var *criaNo(float val, float val2){

	var *newNode;
	
	newNode = malloc(sizeof(var));
	
	if(newNode == NULL){
	
		return NULL;
	
	}
	
	newNode->conteudo[0] = val;
    newNode->conteudo[1] = val2;
	newNode->dir = newNode->esq = NULL;
	return newNode;
	
}

int main(){

	int i,event[100];

	float pd, pac, pr, p_da, p_dr;

	var *raiz = criaNo(1.0,0.2);

	// Calcula a probabilidade do personagem levar dano
	
	pd = raiz->conteudo[1] * 0.1 + (1-raiz->conteudo[1])*0.4;

	raiz->esq	= criaNo(2,pd); // cria nó 2 com probabilidade obtida na linha anterior


	// Calcula a probabilidade de auto cura
	
	pac = raiz->esq->conteudo[1] * 0.25 + (1-raiz->esq->conteudo[1])*0.07;

	raiz->esq->esq	= criaNo(3,pac); // cria nó 2 com probabilidade obtida na linha anterior

	// Calcula a probabilidade da retaliação
	
	pr = raiz->esq->conteudo[1] * 0.3 + (1-raiz->esq->conteudo[1])*0.1;

	raiz->esq->dir	= criaNo(4,pr); // cria nó 2 com probabilidade obtida na linha anterior
        
        for (i = 0;i<=1;i++){
	
		printf("%.2f %.2f %.2f %.2f\n",raiz->conteudo[i], raiz->esq->conteudo[i],raiz->esq->esq->conteudo[i],raiz->esq->dir->conteudo[i]);

	}

	
	// Aplica o teorema de Bayes para obter a P(D | AC) = (P(AC | D) * P(D)) / P(AC)

	p_da = (0.25 * pd)/pac;
	
	// Aplica o teorema de Bayes para obter a P(D | AC) = (P(AC | D) * P(D)) / P(AC)
	
	p_dr = (0.3 * pd)/ pr;

	printf("A probabilidade de dano eh: %f \n",pd);
	printf("A probabilidade da auto cura eh: %f \n",pac);
	printf("A probabilidade da retalhação eh: %f \n",pr);
	printf("A probabilidade do dano dado que o personagem utilizou auto cura eh: %f \n",p_da);
	printf("A probabilidade do dano dado que o personagem utilizou retaliação eh: %f \n",p_dr);

	return 0;

}
